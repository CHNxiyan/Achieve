option(USE_SYSTEM_UVW "Use system libuvw" OFF)

if(USE_SYSTEM_UVW)
    # Special package name from vcpkg
    find_package(unofficial-libuv CONFIG QUIET)
    if(${unofficial-libuv_FOUND})
        add_library(Qv2ray::libuv ALIAS unofficial::libuv::libuv)
    else()
        find_package(LibUV REQUIRED)
        add_library(Qv2ray::libuv ALIAS LibUV::LibUV)
    endif()

    find_package(uvw CONFIG REQUIRED)
    add_library(Qv2ray_uvw INTERFACE)
    target_link_libraries(Qv2ray_uvw INTERFACE uvw::uvw)
    target_link_libraries(Qv2ray_uvw INTERFACE Qv2ray::libuv)
    add_library(Qv2ray::libuvw ALIAS Qv2ray_uvw)
else()
    set(BUILD_UVW_LIBS ON)
    add_subdirectory(${CMAKE_CURRENT_SOURCE_DIR}/3rdparty/uvw)
    add_library(Qv2ray::libuvw ALIAS uvw)

    # BEGIN - the hack to install libuvw as static libraries
    if(NOT BUILD_SHARED_LIBS)
        # See https://github.com/skypjack/uvw/pull/264
        export(EXPORT uvwConfig)
    endif()
    # END  - the hack to install libuvw as static libraries
endif()
