@echo off
setlocal

set DEST="%UserProfile%\.moe-config"

cd %DEST%
git pull

cd %UserProfile%

mklink /J svl %DEST%\svl

if defined AUTOPREP (
    set MOE_MENUS="%DEST%\.moe-menus"
) else (
    set MOE_MENUS="%DEST%\.moe-menus-noautoprep"
)

mklink /H .moe-menus %MOE_MENUS%

pause