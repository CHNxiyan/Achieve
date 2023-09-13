//===----------------------------------------------------------------------===//
//
// Part of the LLVM Project, under the Apache License v2.0 with LLVM Exceptions.
// See https://llvm.org/LICENSE.txt for license information.
// SPDX-License-Identifier: Apache-2.0 WITH LLVM-exception
//
//===----------------------------------------------------------------------===//
//
// WARNING: This test was generated by generate_feature_test_macro_components.py
// and should not be edited manually.
//
// clang-format off

// <string_view>

// Test the feature test macros defined by <string_view>

/*  Constant                           Value
    __cpp_lib_char8_t                  201907L [C++20]
    __cpp_lib_constexpr_string_view    201811L [C++20]
    __cpp_lib_starts_ends_with         201711L [C++20]
    __cpp_lib_string_contains          202011L [C++23]
    __cpp_lib_string_view              201606L [C++17]
                                       201803L [C++20]
*/

#include <string_view>
#include "test_macros.h"

#if TEST_STD_VER < 14

# ifdef __cpp_lib_char8_t
#   error "__cpp_lib_char8_t should not be defined before c++20"
# endif

# ifdef __cpp_lib_constexpr_string_view
#   error "__cpp_lib_constexpr_string_view should not be defined before c++20"
# endif

# ifdef __cpp_lib_starts_ends_with
#   error "__cpp_lib_starts_ends_with should not be defined before c++20"
# endif

# ifdef __cpp_lib_string_contains
#   error "__cpp_lib_string_contains should not be defined before c++23"
# endif

# ifdef __cpp_lib_string_view
#   error "__cpp_lib_string_view should not be defined before c++17"
# endif

#elif TEST_STD_VER == 14

# ifdef __cpp_lib_char8_t
#   error "__cpp_lib_char8_t should not be defined before c++20"
# endif

# ifdef __cpp_lib_constexpr_string_view
#   error "__cpp_lib_constexpr_string_view should not be defined before c++20"
# endif

# ifdef __cpp_lib_starts_ends_with
#   error "__cpp_lib_starts_ends_with should not be defined before c++20"
# endif

# ifdef __cpp_lib_string_contains
#   error "__cpp_lib_string_contains should not be defined before c++23"
# endif

# ifdef __cpp_lib_string_view
#   error "__cpp_lib_string_view should not be defined before c++17"
# endif

#elif TEST_STD_VER == 17

# ifdef __cpp_lib_char8_t
#   error "__cpp_lib_char8_t should not be defined before c++20"
# endif

# ifdef __cpp_lib_constexpr_string_view
#   error "__cpp_lib_constexpr_string_view should not be defined before c++20"
# endif

# ifdef __cpp_lib_starts_ends_with
#   error "__cpp_lib_starts_ends_with should not be defined before c++20"
# endif

# ifdef __cpp_lib_string_contains
#   error "__cpp_lib_string_contains should not be defined before c++23"
# endif

# ifndef __cpp_lib_string_view
#   error "__cpp_lib_string_view should be defined in c++17"
# endif
# if __cpp_lib_string_view != 201606L
#   error "__cpp_lib_string_view should have the value 201606L in c++17"
# endif

#elif TEST_STD_VER == 20

# if defined(__cpp_char8_t)
#   ifndef __cpp_lib_char8_t
#     error "__cpp_lib_char8_t should be defined in c++20"
#   endif
#   if __cpp_lib_char8_t != 201907L
#     error "__cpp_lib_char8_t should have the value 201907L in c++20"
#   endif
# else
#   ifdef __cpp_lib_char8_t
#     error "__cpp_lib_char8_t should not be defined when the requirement 'defined(__cpp_char8_t)' is not met!"
#   endif
# endif

# ifndef __cpp_lib_constexpr_string_view
#   error "__cpp_lib_constexpr_string_view should be defined in c++20"
# endif
# if __cpp_lib_constexpr_string_view != 201811L
#   error "__cpp_lib_constexpr_string_view should have the value 201811L in c++20"
# endif

# ifndef __cpp_lib_starts_ends_with
#   error "__cpp_lib_starts_ends_with should be defined in c++20"
# endif
# if __cpp_lib_starts_ends_with != 201711L
#   error "__cpp_lib_starts_ends_with should have the value 201711L in c++20"
# endif

# ifdef __cpp_lib_string_contains
#   error "__cpp_lib_string_contains should not be defined before c++23"
# endif

# ifndef __cpp_lib_string_view
#   error "__cpp_lib_string_view should be defined in c++20"
# endif
# if __cpp_lib_string_view != 201803L
#   error "__cpp_lib_string_view should have the value 201803L in c++20"
# endif

#elif TEST_STD_VER == 23

# if defined(__cpp_char8_t)
#   ifndef __cpp_lib_char8_t
#     error "__cpp_lib_char8_t should be defined in c++23"
#   endif
#   if __cpp_lib_char8_t != 201907L
#     error "__cpp_lib_char8_t should have the value 201907L in c++23"
#   endif
# else
#   ifdef __cpp_lib_char8_t
#     error "__cpp_lib_char8_t should not be defined when the requirement 'defined(__cpp_char8_t)' is not met!"
#   endif
# endif

# ifndef __cpp_lib_constexpr_string_view
#   error "__cpp_lib_constexpr_string_view should be defined in c++23"
# endif
# if __cpp_lib_constexpr_string_view != 201811L
#   error "__cpp_lib_constexpr_string_view should have the value 201811L in c++23"
# endif

# ifndef __cpp_lib_starts_ends_with
#   error "__cpp_lib_starts_ends_with should be defined in c++23"
# endif
# if __cpp_lib_starts_ends_with != 201711L
#   error "__cpp_lib_starts_ends_with should have the value 201711L in c++23"
# endif

# ifndef __cpp_lib_string_contains
#   error "__cpp_lib_string_contains should be defined in c++23"
# endif
# if __cpp_lib_string_contains != 202011L
#   error "__cpp_lib_string_contains should have the value 202011L in c++23"
# endif

# ifndef __cpp_lib_string_view
#   error "__cpp_lib_string_view should be defined in c++23"
# endif
# if __cpp_lib_string_view != 201803L
#   error "__cpp_lib_string_view should have the value 201803L in c++23"
# endif

#elif TEST_STD_VER > 23

# if defined(__cpp_char8_t)
#   ifndef __cpp_lib_char8_t
#     error "__cpp_lib_char8_t should be defined in c++26"
#   endif
#   if __cpp_lib_char8_t != 201907L
#     error "__cpp_lib_char8_t should have the value 201907L in c++26"
#   endif
# else
#   ifdef __cpp_lib_char8_t
#     error "__cpp_lib_char8_t should not be defined when the requirement 'defined(__cpp_char8_t)' is not met!"
#   endif
# endif

# ifndef __cpp_lib_constexpr_string_view
#   error "__cpp_lib_constexpr_string_view should be defined in c++26"
# endif
# if __cpp_lib_constexpr_string_view != 201811L
#   error "__cpp_lib_constexpr_string_view should have the value 201811L in c++26"
# endif

# ifndef __cpp_lib_starts_ends_with
#   error "__cpp_lib_starts_ends_with should be defined in c++26"
# endif
# if __cpp_lib_starts_ends_with != 201711L
#   error "__cpp_lib_starts_ends_with should have the value 201711L in c++26"
# endif

# ifndef __cpp_lib_string_contains
#   error "__cpp_lib_string_contains should be defined in c++26"
# endif
# if __cpp_lib_string_contains != 202011L
#   error "__cpp_lib_string_contains should have the value 202011L in c++26"
# endif

# ifndef __cpp_lib_string_view
#   error "__cpp_lib_string_view should be defined in c++26"
# endif
# if __cpp_lib_string_view != 201803L
#   error "__cpp_lib_string_view should have the value 201803L in c++26"
# endif

#endif // TEST_STD_VER > 23
