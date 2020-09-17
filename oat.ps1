#!/usr/bin/env pwsh
$cwd = ("$PWD").replace("\", "/")

if ($args[0] -ne "build" -and $args[0] -ne "run") {
    $args = ,"run" + $args
}

chdir "$cwd"
& "$PSScriptRoot\kast_start.exe" $args