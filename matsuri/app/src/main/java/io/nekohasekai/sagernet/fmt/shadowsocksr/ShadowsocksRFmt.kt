/******************************************************************************
 *                                                                            *
 * Copyright (C) 2021 by nekohasekai <contact-sagernet@sekai.icu>             *
 *                                                                            *
 * This program is free software: you can redistribute it and/or modify       *
 * it under the terms of the GNU General Public License as published by       *
 * the Free Software Foundation, either version 3 of the License, or          *
 *  (at your option) any later version.                                       *
 *                                                                            *
 * This program is distributed in the hope that it will be useful,            *
 * but WITHOUT ANY WARRANTY; without even the implied warranty of             *
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the              *
 * GNU General Public License for more details.                               *
 *                                                                            *
 * You should have received a copy of the GNU General Public License          *
 * along with this program. If not, see <http://www.gnu.org/licenses/>.       *
 *                                                                            *
 ******************************************************************************/

package io.nekohasekai.sagernet.fmt.shadowsocksr

import io.nekohasekai.sagernet.ktx.applyDefaultValues
import io.nekohasekai.sagernet.ktx.decodeBase64UrlSafe
import moe.matsuri.nya.utils.Util
import okhttp3.HttpUrl.Companion.toHttpUrl
import org.json.JSONObject
import java.util.*

fun parseShadowsocksR(url: String): ShadowsocksRBean {

    val params = url.substringAfter("ssr://").decodeBase64UrlSafe().split(":")

    val bean = ShadowsocksRBean().apply {
        serverAddress = params[0]
        serverPort = params[1].toInt()
        protocol = params[2]
        method = params[3]
        obfs = params[4]
        password = params[5].substringBefore("/").decodeBase64UrlSafe()
    }

    val httpUrl = ("https://localhost" + params[5].substringAfter("/")).toHttpUrl()

    runCatching {
        bean.obfsParam = httpUrl.queryParameter("obfsparam")!!.decodeBase64UrlSafe()
    }
    runCatching {
        bean.protocolParam = httpUrl.queryParameter("protoparam")!!.decodeBase64UrlSafe()
    }

    val remarks = httpUrl.queryParameter("remarks")
    if (!remarks.isNullOrBlank()) {
        bean.name = remarks.decodeBase64UrlSafe()
    }

    return bean

}

fun ShadowsocksRBean.toUri(): String {

    return "ssr://" + Util.b64EncodeUrlSafe(
        "%s:%d:%s:%s:%s:%s/?obfsparam=%s&protoparam=%s&remarks=%s".format(
            Locale.ENGLISH,
            serverAddress,
            serverPort,
            protocol,
            method,
            obfs,
            Util.b64EncodeUrlSafe("%s".format(Locale.ENGLISH, password)),
            Util.b64EncodeUrlSafe("%s".format(Locale.ENGLISH, obfsParam)),
            Util.b64EncodeUrlSafe("%s".format(Locale.ENGLISH, protocolParam)),
            Util.b64EncodeUrlSafe(
                "%s".format(
                    Locale.ENGLISH, name ?: ""
                )
            )
        )
    )
}


fun JSONObject.parseShadowsocksR(): ShadowsocksRBean {
    return ShadowsocksRBean().applyDefaultValues().apply {
        serverAddress = optString("server", serverAddress)
        serverPort = optInt("server_port", serverPort)
        method = optString("method", method)
        password = optString("password", password)
        protocol = optString("protocol", protocol)
        protocolParam = optString("protocol_param", protocolParam)
        obfs = optString("obfs", obfs)
        obfsParam = optString("obfs_param", obfsParam)
        name = optString("remarks", name)
    }
}
