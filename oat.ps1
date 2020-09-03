#!/usr/bin/env pwsh
$cwd = ("$PWD").replace("\", "/")
& "$PSScriptRoot\oat_start.exe" $args -cwd="$cwd"