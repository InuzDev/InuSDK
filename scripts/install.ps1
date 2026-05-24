$ErrorActionPreference = "Stop"

# Define the repository, installation directoy and the binary path.
$repo = "InuzDev/InuSDK"
$shimsDir = "$env:USERPROFILE\.inusdk\shims"
$binPath = "$shimsDir\inusdk.exe"

Write-Host ""
Write-Host "Installing InuSDK. . ." -ForegroundColor Cyan

# Check the architecture of the machine
$arch = if ([System.Environment]::Is64BitOperatingsystem)
{ "amd64"
} else
{ "386"
}
Write-Host "Detected: Windows/$arch"

# Fetch the latest release from the repository release page
Write-Host "Fetching latest release"
$release = Invoke-RestMethod -Uri "https://api.github.com/repos/$repo/releases/latest"
$version = $release.tag_name

$asset = $release.assets | Where-Object {
   $_.name -like "*windows*$arch*.zip" -and $_.name -notlike "*checksums*"
}

# If there wasn't any available version for the user's machine.
if (-not $asset)
{
   Write-Host "No binary found for windows/$arch" -ForegroundColor Red
   exit 1
}

# Create the directory where the CLI will be located
New-Item -ItemType Directory -Force -Path $shimsDir | Out-Null

# Download the binary
Write-Host "Downloading $version for windows/$arch. . ."
$tempFile = "$env:TEMP\inusdk.zip"
Invoke-WebRequest -Uri $asset.browser_download_url -OutFile $tempFile

# Extract the zip file
Write-Host "Extracting files . . ."
Expand-Archive -Path $tempFile -DestinationPath "$env:TEMP\inusdk_extract" -Force
Copy-Item "$env:TEMP\inusdk_extract\inusdk.exe" $binPath -Force

# Clean the tempfile by deleting it
Remove-Item $tempFile -Force
Remove-Item "$env:TEMP\inusdk_extract" -Recurse -Force

# Add the CLI application to PATH (It will check if it wasn't added manually)
$currentPath = [System.Environment]::GetEnvironmentVariable("PATH", "User")
if ($currentPath -notlike "*$shimsDir*")
{
   Write-Host "Adding to PATH"
   [System.Environment]::SetEnvironmentVariable("PATH", "$shimsDir;$currentPath", "User")
   Write-Host "Added to PATH succesfully" -ForegroundColor Green
} else
{
   Write-Host "The application is already in PATH, if there is any error. Add it manually" -ForegroundColor Yellow
}

# Cleanup the temp files

Remove-Item $tempFile -ErrorAction SilentlyContinue -Force

Write-Host ""
Write-Host "InuSDK $version installed succesfully" -ForegroundColor Magenta
Write-Host "Restart your terminal and run inusdk" -ForegroundColor Cyan
Write-Host ""
