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

package io.nekohasekai.sagernet.fmt.shadowsocks

import com.github.shadowsocks.plugin.PluginConfiguration
import com.github.shadowsocks.plugin.PluginOptions
import io.nekohasekai.sagernet.ktx.*
import moe.matsuri.nya.utils.Util
import okhttp3.HttpUrl.Companion.toHttpUrlOrNull
import org.json.JSONObject

fun PluginConfiguration.fixInvalidParams() {

    if (selected.contains("v2ray") && selected != "v2ray-plugin") {

        pluginsOptions["v2ray-plugin"] = getOptions().apply { id = "v2ray-plugin" }
        pluginsOptions.remove(selected)
        selected = "v2ray-plugin"

        // resolve v2ray plugin

    }

    if (selected.contains("obfs") && selected != "obfs-local") {

        pluginsOptions["obfs-local"] = getOptions().apply { id = "obfs-local" }
        pluginsOptions.remove(selected)
        selected = "obfs-local"

        // resolve clash obfs

    }

    if (selected == "obfs-local") {
        val options = pluginsOptions["obfs-local"]
        if (options != null) {
            if (options.containsKey("mode")) {
                options["obfs"] = options["mode"]
                options.remove("mode")
            }
            if (options.containsKey("host")) {
                options["obfs-host"] = options["host"]
                options.remove("host")
            }
        }
    }

}

fun ShadowsocksBean.fixInvalidParams() {
    if (method == "plain") method = "none"
    plugin = PluginConfiguration(plugin).apply { fixInvalidParams() }.toString()

}

fun parseShadowsocks(url: String): ShadowsocksBean {

    if (url.substringBefore("#").contains("@")) {

        var link = url.replace("ss://", "https://").toHttpUrlOrNull() ?: error(
            "invalid ss-android link $url"
        )

        if (link.username.isBlank()) { // fix justmysocks's shit link

            link = (("https://" + url.substringAfter("ss://")
                .substringBefore("#")
                .decodeBase64UrlSafe()).toHttpUrlOrNull() ?: error(
                "invalid jms link $url"
            )).newBuilder().fragment(url.substringAfter("#")).build()
        }

        // ss-android style

        if (link.password.isNotBlank()) {

            return ShadowsocksBean().apply {

                serverAddress = link.host
                serverPort = link.port
                method = link.username
                password = link.password
                plugin = link.queryParameter("plugin") ?: ""
                name = link.fragment

                fixInvalidParams()

            }

        }

        val methodAndPswd = link.username.decodeBase64UrlSafe()

        return ShadowsocksBean().apply {

            serverAddress = link.host
            serverPort = link.port
            method = methodAndPswd.substringBefore(":")
            password = methodAndPswd.substringAfter(":")
            plugin = link.queryParameter("plugin") ?: ""
            name = link.fragment

            fixInvalidParams()

        }

    } else {
        // v2rayN style

        var v2Url = url

        if (v2Url.contains("#")) v2Url = v2Url.substringBefore("#")

        val link = ("https://" + v2Url.substringAfter("ss://")
            .decodeBase64UrlSafe()).toHttpUrlOrNull() ?: error("invalid v2rayN link $url")

        return ShadowsocksBean().apply {

            serverAddress = link.host
            serverPort = link.port
            method = link.username
            password = link.password
            plugin = ""
            val remarks = url.substringAfter("#").unUrlSafe()
            if (remarks.isNotBlank()) name = remarks
            fixInvalidParams()

        }

    }

}

fun ShadowsocksBean.toUri(): String {

    val builder = linkBuilder().username(Util.b64EncodeUrlSafe("$method:$password"))
        .host(serverAddress)
        .port(serverPort)

    if (plugin.isNotBlank()) {
        builder.addQueryParameter("plugin", plugin)
    }

    if (name.isNotBlank()) {
        builder.encodedFragment(name.urlSafe())
    }

    return builder.toLink("ss").replace("$serverPort/", "$serverPort")

}

fun JSONObject.parseShadowsocks(): ShadowsocksBean {
    return ShadowsocksBean().apply {
        var pluginStr = ""
        val pId = getStr("plugin")
        if (!pId.isNullOrBlank()) {
            val plugin = PluginOptions(pId, getStr("plugin_opts"))
            pluginStr = plugin.toString(false)
        }

        serverAddress = getStr("server")
        serverPort = getIntNya("server_port")
        password = getStr("password")
        method = getStr("method")
        plugin = pluginStr
        name = optString("remarks", "")

        fixInvalidParams()
    }
}