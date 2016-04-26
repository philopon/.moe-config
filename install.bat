@echo off
setlocal

set DEST="%UserProfile%\.moe-config\"

if exist %DEST% (
    cd %DEST%
    git pull
) else (
    git clone "https://github.com/philopon/.moe-config.git" %DEST%
    cd %DEST%
)

cd %UserProfile%

mklink /J svl %DEST%\svl

if defined AUTOPREP (
    set MOE_MENUS="%DEST%\.moe-menus"
) else (
    set MOE_MENUS="%DEST%\.moe-menus-noautoprep"
)

mklink /H .moe-menus %MOE_MENUS%

set /p dummy="done!"
