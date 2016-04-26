@echo off
setlocal

cd %~dp0
git pull

cd %HOME%

mklink /J svl %~dp0\svl 2> nul
mklink /H .moe-menus %~dp0\.moe-menus 2> nul
