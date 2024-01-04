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

# get version from arguments or use the latest
$VERSION = $args[0] ?? "latest"
$TOOL_NAME = "aqua"
$REPO = "aquaproj/${TOOL_NAME}"
$TOOL_DIR = $REPO.Replace("/", "-")
$TOOL_HOME="$env:LOCALAPPDATA\$TOOL_DIR"

# provides a way to get the architecture of the current process
# while allowing for easy mapping to the architecture names used by the GitHub API
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
function get_assets() {
    $url = "https://api.github.com/repos/$REPO/releases/$VERSION"
    
    # return assets
    (Invoke-RestMethod -Method GET -Uri $url).assets
}

# Get the url for the windows release
# https://developer.github.com/v3/repos/releases/#get-a-single-release-asset
function get_windows_release_url() {
    $assets = get_assets
    $arch = get_architecture

    # get windows release
    $windows_release = $assets | Where-Object name -EQ "${TOOL_NAME}_windows_$arch.zip"

    # return browser_download_url
    return $windows_release.browser_download_url
}

# Download the windows release
# https://developer.github.com/v3/repos/releases/#get-a-single-release-asset
# Creates a temporary directory, downloads the release, unzips it, and copies the files to $TOOL_HOME
function download_windows_release() {
    $url = get_windows_release_url
    echo "Downloading $url to $output"
    mkdir -f $TOOL_HOME

    # make to tmp directory
    $tempFolderPath = Join-Path $Env:Temp $(New-Guid); New-Item -Type Directory -Path $tempFolderPath | Out-Null

    # download file
    (new-object System.Net.WebClient).DownloadFile("$url", "$tempFolderPath\tool.zip")

    # unzip 
    Add-Type -AssemblyName System.IO.Compression.FileSystem
    [System.IO.Compression.ZipFile]::ExtractToDirectory("$tempFolderPath\tool.zip", "$tempFolderPath")

    # remove zip 
    Remove-Item "$tempFolderPath\tool.zip"

    # copy files to $TOOL_HOME merging directories
    Copy-Item "$tempFolderPath\*" "$TOOL_HOME" -Recurse -Force
}

download_windows_release

echo ""
echo "Installed $(& $TOOL_HOME\$TOOL_NAME --version)"
echo ""
echo "Installation complete."
echo "Add $TOOL_NAME to your path by adding the following line to your ~/Documents/Powershell/Microsoft.PowerShell_profile.ps1 and open a new terminal:"
echo "  $(-join('$env:PATH+=";', $TOOL_HOME, '"'))"
