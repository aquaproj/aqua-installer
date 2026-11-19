# 
# Download and install the latest version of aqua
#
# Usage:
#
#   # install the latest version
#   iwr https://raw.githubusercontent.com/aquaproj/aqua-installer/master/aqua-installer.ps1 | iex
#   
#   # install a specific version
#   & ([scriptblock]::Create((iwr https://raw.githubusercontent.com/aquaproj/aqua-installer/master/aqua-installer.ps1))) -param1 v2.2.0
#

param([String]$version) 


# get version from arguments or use the latest
$VERSION = $version
$BOOTSTRAP_VERSION = "v2.21.3"
$TOOL_NAME = "aqua"
$REPO = "aquaproj/${TOOL_NAME}"
$TOOL_DIR = $REPO.Replace("/", "-")
$TOOL_HOME="$env:LOCALAPPDATA\$TOOL_DIR"


$hashes = @"
489d8e279d437f4a8710ab6807420358499c1c33294fb58f6a97b18899250cb4  aqua_darwin_amd64.tar.gz
30003f594d92f3b9d547256af9344ccb1ee8e4b1f8b739d905646397672ffa55  aqua_darwin_arm64.tar.gz
8c4ba7f95563c273b44d90045108e8f8551a6beef74690161baf2bd0299ab0e3  aqua_linux_amd64.tar.gz
d55d0c5f6d7701e9643c72b373265460fd52187c1e70fefc3599ce346065a20c  aqua_linux_arm64.tar.gz
e333384833467ec6f632ca2e81ae63b02d3c02b750fc5b5d830d791b85f78aa4  aqua_windows_amd64.zip
7b7ffcb90bdf5b8e26a8c9af2069612fd31e96585f134c93ea22119f5afbe40e  aqua_windows_arm64.zip
"@ -split "`n"

function log() {
    $message = $args[0]
    $level = $args[1] ?? "INFO"

    write-host "[$level] $message"
}

function get_env_paths() {
    param([String]$path)
    
    # return the command to add the given path to the path environment variable
    $paths=@(
        "$path",
        # join the path using the path separator for the current platform
        [IO.Path]::Combine($path, "bin"),
        [IO.Path]::Combine($path, "bat")
    ) -join ";"
    
    return $paths
}

function get_env_path_command() {
    param([String]$path)

    # return the command to add the given path to the path environment variable
    return '$env:PATH=$env:PATH;' + $(get_env_paths $path)
}

function get_file_hash() {
    param([String]$filepath)
    $hash = get-filehash -Path "$filepath" -Algorithm SHA256
    return $hash
}

# Verify that the file hash matches the expected hash
# https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.utility/get-filehash?view=powershell-7.1
function verify_release_hash() {
    param([String]$filepath)
    log "Verifying $filepath"

    # get the filename from the path
    $file = "${filepath}".Split("\")[-1]

    log "Verifying $file hash"

    #loop through the hashes array and find the hash for the given file
    $hash = $hashes | Where-Object { $_.Contains($file) } | ForEach-Object { $_.Split(" ")[0] }
    
    if ($hash -eq $null) {
        throw "${file} not valid name."
    }

    $computed_hash = (Get-FileHash "$filepath").Hash

    if ($hash -ne $computed_hash) {
        throw "Hash mismatch: $hash != $computed_hash"
    }

    log "Hash verified"
}

# provides a way to get the architecture of the current process
# while allowing for easy mapping to the architecture names used by the GitHub API
# https://developer.github.com/v3/repos/releases/#get-a-release-by-tag-name
function get_architecture() {
    $arch = $env:PROCESSOR_ARCHITECTURE

    switch ($arch) {
        "x86" { return "x86" }
        "AMD64" { return "amd64" }
        "arm64" { return "arm64" }
        default { return "unknown" }
    }

    return "unknown"
}

# Get a list of assets for the given release
# https://developer.github.com/v3/repos/releases/#get-a-release-by-tag-name
function get_assets_for_tag() {
    param([String]$tag)

    $url = "https://api.github.com/repos/$REPO/releases/tags/$tag"

    log "Getting assets for $url"
    
    # return assets
    (Invoke-RestMethod -Method GET -Uri $url).assets
}

# Get the url for the windows release
# https://developer.github.com/v3/repos/releases/#get-a-single-release-asset
function get_windows_release_url() {
    param([String]$version)
    $assets = get_assets_for_tag $version
    $arch = get_architecture

    # get windows release
    $windows_release = $assets | Where-Object name -EQ "${TOOL_NAME}_windows_$arch.zip"

    # return browser_download_url
    return $windows_release.browser_download_url
}

# Download the windows release
# https://developer.github.com/v3/repos/releases/#get-a-single-release-asset
# Creates a temporary directory, downloads the release
function download_windows_release() {
    param([String]$version)
    $url = get_windows_release_url $version
    mkdir -f $TOOL_HOME | Out-Null
    
    # make to tmp directory
    $tempFolderPath = Join-Path $Env:Temp $(New-Guid); 
    
    $created = New-Item -Type Directory -Path $tempFolderPath | Out-Null
    
    log "Downloading $url"

    # get the filename from the url
    $filename = $url.Split("/")[-1]
    $destination = "${tempFolderPath}\${filename}"

    # download file
    (new-object System.Net.WebClient).DownloadFile("$url", $destination) | Out-Null
    
    # return the path to the downloaded file
    return "${destination}"
}

# Install the windows release
#  unzips it, and copies the files to $TOOL_HOME
function install_windows_release() {
    param([String]$filepath)

    log "Installing $TOOL_HOME"

    # find the parent folder of the zip
    $parent = Split-Path $filepath

    # unzip 
    Add-Type -AssemblyName System.IO.Compression.FileSystem | Out-Null
    [System.IO.Compression.ZipFile]::ExtractToDirectory($filepath, "$parent") | Out-Null

    # remove zip 
    Remove-Item $filepath | Out-Null

    # for each item in the parent folder, move it to $TOOL_HOME overwriting existing files
    ls "$parent\*" | ForEach-Object { 
        $item = $_
        $destination = "$TOOL_HOME\$($item.Name)"
        Remove-Item $destination -Force -ErrorAction SilentlyContinue | Out-Null
        log "Copying $($item.Name) to $destination"
        Copy-Item $item $destination -Force | Out-Null
     }

    # ensure TOOL_HOME/bin and TOOL_HOME/bat are in the path
    # iex $(get_set_env_path_command $TOOL_HOME)

    ls $TOOL_HOME
    write-host " "
}

# Runs the aqua upgrade command to upgrade to the given version
function upgrade_to_version() {
    param([String]$version=$VERSION)

    # ensure the TOOL_HOME/bin and TOOL_HOME/bat are in the path
    $env:PATH="$env:PATH;$(get_env_paths $path)"

    # if version is not supplied, just use the upgrade command without args
    if ($version -eq $null -or $version -eq "") {
        $version = "latest"
    }

    log ""
    log "Upgrading to $version"
    log ""
    
    log "upgrading ... $(& $TOOL_HOME\$TOOL_NAME update-aqua $version)"

    return $outcome
}

# download the version that matches the hashes above
$release_zip = download_windows_release "$BOOTSTRAP_VERSION"

# # verify the hash of the downloaded file
verify_release_hash "$release_zip"

# # install the release
install_windows_release "$release_zip"

# # install the specified version or the latest version
upgrade_to_version "$VERSION"

echo ""
echo "Installed $(& $TOOL_HOME\$TOOL_NAME --version)"
echo ""
echo "Installation complete."
echo ""
echo "Add $TOOL_NAME to your path by following these steps:"
echo ""
echo "  1. Open '$profile' in your favorite editor"
echo "  2. Add the following line to the end of the file:"
echo ""
echo "    $(get_env_path_command $TOOL_HOME)"
echo ""
echo "  3. Save and close the file"
echo "  4. Restart your terminal"

