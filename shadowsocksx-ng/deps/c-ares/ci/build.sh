#!/bin/sh
set -e

OS=""
if [ "$TRAVIS_OS_NAME" != "" ]; then
    OS="$TRAVIS_OS_NAME"
elif [ "$CIRRUS_OS" != "" ]; then
    OS="$CIRRUS_OS"
fi

if [ "$BUILD_TYPE" != "cmake" -a "$BUILD_TYPE" != "valgrind" -a "$BUILD_TYPE" != "ios-cmake" ]; then
    autoreconf -fi
    mkdir atoolsbld
    cd atoolsbld
    $SCAN_WRAP ../configure --disable-symbol-hiding --enable-expose-statics --enable-maintainer-mode --enable-debug $CONFIG_OPTS
    $SCAN_WRAP make
else
    if [ "$BUILD_TYPE" = "ios-cmake" ]; then
        mkdir cmakebld
        cd cmakebld
        cmake \
          -DCMAKE_BUILD_TYPE=DEBUG                 \
          -DCARES_STATIC=ON                        \
          -DCARES_STATIC_PIC=ON                    \
          -DCARES_BUILD_TESTS=OFF                  \
          -DCMAKE_C_FLAGS=$CFLAGS                  \
          -DCMAKE_CXX_FLAGS=$CXXFLAGS              \
          -DCMAKE_OSX_SYSROOT=$SYSROOT             \
          -DCMAKE_OSX_ARCHITECTURES=$ARCHITECTURES \
          ..
        make
    else
        # Use cmake for valgrind to prevent libtool script wrapping of tests that interfere with valgrind
        mkdir cmakebld
        cd cmakebld
        cmake -DCMAKE_BUILD_TYPE=DEBUG -DCARES_STATIC=ON -DCARES_STATIC_PIC=ON -DCARES_BUILD_TESTS=ON ..
        make
    fi
fi
