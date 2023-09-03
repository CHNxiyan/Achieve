#!/bin/bash
# 打包,发布

APP_NAME="V2rayU"
INFOPLIST_FILE="Info.plist"
BASE_DIR=$HOME/swift/V2rayU
BUILD_DIR=${BASE_DIR}/Build
V2rayU_ARCHIVE=${BUILD_DIR}/V2rayU.xcarchive
V2rayU_RELEASE=${BUILD_DIR}/release
APP_Version=$(sed -n '/MARKETING_VERSION/{s/MARKETING_VERSION = //;s/;//;s/^[[:space:]]*//;p;q;}' ../V2rayU.xcodeproj/project.pbxproj)
DMG_FINAL="${APP_NAME}.dmg"
APP_TITLE="${APP_NAME} - V${APP_Version}"
AppCastDir=$HOME/swift/appcast

function updatePlistVersion() {
    buildString=$(/usr/libexec/PlistBuddy -c "Print CFBundleShortVersionString" "${BASE_DIR}/V2rayU/${INFOPLIST_FILE}")
    /usr/libexec/PlistBuddy -c "Set :CFBundleVersion $buildString" "${BASE_DIR}/V2rayU/${INFOPLIST_FILE}"
}

function build() {
    echo "Building V2rayU."${APP_Version}
    echo "Cleaning up old archive & app..."
    rm -rf ${V2rayU_ARCHIVE} ${V2rayU_RELEASE}

    echo "Building archive... please wait a minute"
    xcodebuild -workspace ${BASE_DIR}/V2rayU.xcworkspace -config Release -scheme V2rayU -archivePath ${V2rayU_ARCHIVE} archive

    echo "Exporting archive..."
    xcodebuild -archivePath ${V2rayU_ARCHIVE} -exportArchive -exportPath ${V2rayU_RELEASE} -exportOptionsPlist ./build.plist

    echo "Cleaning up archive..."
    rm -rf ${V2rayU_ARCHIVE}

    chmod -R 755 "${V2rayU_RELEASE}/${APP_NAME}.app/Contents/Resources/v2ray-core"
    chmod -R 755 "${V2rayU_RELEASE}/${APP_NAME}.app/Contents/Resources/unzip.sh"
}

function createDmg() {
    umount "/Volumes/${APP_NAME}"

    ############# 1 #############
    APP_PATH="${V2rayU_RELEASE}/${APP_NAME}.app"
    DMG_BACKGROUND_IMG="dmg-bg@2x.png"

    DMG_TMP="${APP_NAME}-temp.dmg"

    # 清理文件夹
    echo "createDmg start."
    rm -rf "${DMG_TMP}" "${DMG_FINAL}"
    # 创建文件夹，拷贝，计算
    SIZE=`du -sh "${APP_PATH}" | sed 's/\([0-9\.]*\)M\(.*\)/\1/'`
    SIZE=`echo "${SIZE} + 1.0" | bc | awk '{print int($1+0.5)}'`
    # 容错处理
    if [ $? -ne 0 ]; then
       echo "Error: Cannot compute size of staging dir"
       exit
    fi
    # 创建临时dmg文件
    hdiutil create -srcfolder "${APP_PATH}" -volname "${APP_NAME}" -fs HFS+ \
          -fsargs "-c c=64,a=16,e=16" -format UDRW -size ${SIZE}M "${DMG_TMP}"
    echo "Created DMG: ${DMG_TMP}"

    ############# 2 #############
    DEVICE=$(hdiutil attach -readwrite -noverify "${DMG_TMP}"| egrep '^/dev/' | sed 1q | awk '{print $1}')

    # 拷贝背景图片
    mkdir /Volumes/"${APP_NAME}"/.background
    cp "${BUILD_DIR}/${DMG_BACKGROUND_IMG}" /Volumes/"${APP_NAME}"/.background/
    # 使用applescript设置一系列的窗口属性
    echo '
       tell application "Finder"
         tell disk "'${APP_NAME}'"
               open
               set current view of container window to icon view
               set toolbar visible of container window to false
               set statusbar visible of container window to false
               set the bounds of container window to {0, 0, 560, 297}
               set viewOptions to the icon view options of container window
               set arrangement of viewOptions to not arranged
               set icon size of viewOptions to 80
               set background picture of viewOptions to file ".background:'${DMG_BACKGROUND_IMG}'"
               make new alias file at container window to POSIX file "/Applications" with properties {name:"Applications"}
               delay 1
               set position of item "'${APP_NAME}'.app" of container window to {120, 120}
               set position of item "Applications" of container window to {380, 120}
               close
               open
               update without registering applications
               delay 2
         end tell
       end tell
    ' | osascript

    sync
    # 卸载
    hdiutil detach "${DEVICE}"

    ############# 3 #############
    echo "Creating compressed image"
    hdiutil convert "${DMG_TMP}" -format UDZO -imagekey zlib-level=9 -o "${DMG_FINAL}"

    # appcast sign update
    ${AppCastDir}/bin/sign_update ${DMG_FINAL}

    umount "/Volumes/${APP_NAME}"
}

function generateAppcast() {
    echo "generate appcast"
    description=$1
    if [[ -z "$description" ]]; then
        description="bug fix"
    fi
    downloadUrl="https://github.com/yanue/V2rayU/releases/download/${APP_Version}/V2rayU-64.dmg"
    # https://github.com/c9s/appcast.git
    ${AppCastDir}/appcast -append\
        -dsaSignature="PW8pDnr5VZkmC93gZjUDlHI8gkJSspPoDU3DdhsMkps"\
        -title="${APP_TITLE}"\
        -description="${description}"\
        -file "${DMG_FINAL}"\
        -url "${downloadUrl}"\
        -version "${APP_Version}"\
        -versionShortString="${APP_Version}"\
        ${AppCastDir}/appcast.xml
}

function pushRelease() {
    description=$1
    if [[ -z "$description" ]]; then
        description="bug fix"
    fi

    echo "github-release tag"
        
    ${AppCastDir}/github-release release\
        --security-token ="f4ff9dc62cdf998cd57f22be811c6df6e2a58050"\
        --user "yanue"\
        --repo "${APP_NAME}"\
        --tag "${APP_Version}"\
        --name "${APP_Version}"\
        --description "${description}"\

    echo "github-release upload"
    ${AppCastDir}/github-release upload\
        --security-token "f4ff9dc62cdf998cd57f22be811c6df6e2a58050"\
        --user "yanue"\
        --repo "${APP_NAME}"\
        --tag "${APP_Version}"\
        --name "${DMG_FINAL}"\
        --file "${DMG_FINAL}"\

    echo "github-release done."
}

function commit() {
    cd ${AppCastDir}
      # commit
    echo "commit push"
    git add appcast.xml
    git commit -a -m "update version: ${APP_Version}"
    git push

    cd ${BUILD_DIR}
    git commit -a -m "update version: ${APP_Version}"
    git push
}

function downloadV2ray() {
    DMG_FINAL="${APP_NAME}-64.dmg"
    rm -fr ${DMG_FINAL} ${V2rayU_RELEASE}
    
    echo "正在查询最新版v2ray-64 ..."
    rm -fr v2ray-core
    tag='v1.5.9'
    echo "v2ray-core version: ${tag}"
    url="https://github.com/XTLS/Xray-core/releases/download/v1.5.9/Xray-macos-64.zip"
    echo "正在下载最新版v2ray: ${tag}"
    curl -Lo Xray-macos-64.zip ${url}

    unzip -o Xray-macos-64.zip -d v2ray-core
    \cp v2ray-core/xray v2ray-core/v2ray
}

function downloadV2rayArm() {
    DMG_FINAL="${APP_NAME}-arm64.dmg"
    rm -fr ${DMG_FINAL} ${V2rayU_RELEASE}

    echo "正在查询最新版v2ray-arm64 ..."
    rm -fr v2ray-core
    tag='v1.5.9'
    echo "v2ray-core version: ${tag}"
    url="https://github.com/XTLS/Xray-core/releases/download/v1.5.9/Xray-macos-arm64-v8a.zip"
    echo "正在下载最新版v2ray: ${tag}"
    curl -Lo Xray-macos-arm64-v8a.zip ${url}

    unzip -o Xray-macos-arm64-v8a.zip -d v2ray-core
    \cp v2ray-core/xray v2ray-core/v2ray
}

function createDmgByAppdmg() {
#    umount "/Volumes/${APP_NAME}"

#    rm -rf ${BUILD_DIR}/${APP_NAME}.app ${BUILD_DIR}/${DMG_FINAL}
#    \cp -Rf "${V2rayU_RELEASE}/${APP_NAME}.app" "${BUILD_DIR}/${APP_NAME}.app"

    rm -f  ${BUILD_DIR}/${DMG_FINAL}
    # https://github.com/LinusU/node-appdmg
    # npm install -g appdmg
    echo ${BUILD_DIR}/appdmg.json
    appdmg appdmg.json ${DMG_FINAL}

    # appcast sign update
    ${AppCastDir}/bin/sign_update ${DMG_FINAL}

#    umount "/Volumes/${APP_NAME}"
}

function makeDmg() {
    echo "正在打包版本: V"${APP_Version}
    read -n1 -r -p "请确认版本号是否正确 [Y/N]? " answer
    case ${answer} in
    Y | y ) echo
    echo "你选择了Y";;
    N | n ) echo
    echo ""
    echo "OK, goodbye"
    exit;;
    *)
    echo ""
    echo "请输入Y|N"
    exit;;
    esac
    echo "请选择build的版本 :"
    options=("64" "arm64")
    select target in "${options[@]}"
    do
        case $target in
        "64")
            echo "你选择了: 64"
#            downloadV2ray
            cd release/
            break
            ;;
        "arm64")
            echo "你选择了: arm64"
#            downloadV2rayArm
            break
            ;;
        *) echo "请选择";;
        esac
    done

#    updatePlistVersion
#    build
    createDmgByAppdmg
}





echo 'done'
