@echo off
setlocal

cd %~dp0
git pull

cd %HOME%

mklink /J svl %~dp0\svl

if "%AUTOPREP%" == "1" (
    set MOE_MENUS="%~dp0\.moe-menus"
) else (
    set MOE_MENUS="%~dp0\.moe-menus-noautoprep"
)

mklink /H .moe-menus %MOE_MENUS%
