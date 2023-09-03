/*
 * Copyright 2021 Squircle IDE contributors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package com.blacksquircle.ui.language.plaintext.styler

import com.blacksquircle.ui.language.base.model.SyntaxScheme
import com.blacksquircle.ui.language.base.span.SyntaxHighlightSpan
import com.blacksquircle.ui.language.base.styler.LanguageStyler
import com.blacksquircle.ui.language.base.utils.StylingResult

class PlainTextStyler private constructor() : LanguageStyler {

    companion object {

        private var plainTextStyler: PlainTextStyler? = null

        fun getInstance(): PlainTextStyler {
            return plainTextStyler ?: PlainTextStyler().also {
                plainTextStyler = it
            }
        }
    }

    override fun execute(
        sourceCode: String,
        syntaxScheme: SyntaxScheme
    ): List<SyntaxHighlightSpan> = emptyList()

    override fun enqueue(
        sourceCode: String,
        syntaxScheme: SyntaxScheme,
        stylingResult: StylingResult
    ) = stylingResult(execute(sourceCode, syntaxScheme))

    override fun cancel() = Unit
}