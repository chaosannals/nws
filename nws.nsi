; ------------------------------
Name "NginxWindowsServiceSetup"
OutFile "nws-setup.exe"
RequestExecutionLevel admin
Unicode True
SetCompressor /SOLID lzma
SetCompress force
InstallDir $PROGRAMFILES\Nws

; ------------------------------

Page components
Page directory
Page instfiles

UninstPage uninstConfirm
UninstPage instfiles

; ------------------------------
; 安装
Section "Install"
    SectionIn RO
    SetOutPath $INSTDIR

    File "/oname=nws.exe" "nws.exe"
    File "/oname=seelog.xml" "seelog.xml"
    NSISdl::download "http://nginx.org/download/nginx-1.21.3.zip" "nginx-1.21.3.zip"
    nsExec::Exec "$INSTDIR\nws install"
    Delete "nginx-1.21.3.zip"

    WriteUninstaller "$INSTDIR\uninstall.exe"
SectionEnd

; 卸载
Section "Uninstall"
    nsExec::Exec "$INSTDIR\nws uninstall"
    Delete "$INSTDIR\nws.exe"
    Delete "$INSTDIR\seelog.xml"
    Delete "$INSTDIR\uninstall.exe"
    RMDir /r "$INSTDIR\nginx"

    RMDir "$INSTDIR"
SectionEnd
