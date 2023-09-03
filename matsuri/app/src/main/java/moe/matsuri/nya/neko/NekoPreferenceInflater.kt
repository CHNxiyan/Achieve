package moe.matsuri.nya.neko

import androidx.preference.Preference
import androidx.preference.PreferenceScreen
import androidx.preference.SwitchPreference
import com.takisoft.preferencex.EditTextPreference
import com.takisoft.preferencex.PreferenceCategory
import com.takisoft.preferencex.SimpleMenuPreference
import io.nekohasekai.sagernet.database.preference.EditTextPreferenceModifiers
import io.nekohasekai.sagernet.ktx.forEach
import io.nekohasekai.sagernet.ktx.getStr
import io.nekohasekai.sagernet.ui.profile.ProfileSettingsActivity
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import moe.matsuri.nya.utils.getDrawableByName
import org.json.JSONArray
import org.json.JSONObject

object NekoPreferenceInflater {
    suspend fun inflate(pref: JSONArray, preferenceScreen: PreferenceScreen) =
        withContext(Dispatchers.Main) {
            val context = preferenceScreen.context
            pref.forEach { _, category ->
                category as JSONObject

                val preferenceCategory = PreferenceCategory(context)
                preferenceScreen.addPreference(preferenceCategory)

                category.getStr("key")?.apply { preferenceCategory.key = this }
                category.getStr("title")?.apply { preferenceCategory.title = this }

                category.optJSONArray("preferences")?.forEach { _, any ->
                    if (any is JSONObject) {
                        lateinit var p: Preference
                        // Create Preference
                        when (any.getStr("type")) {
                            "EditTextPreference" -> {
                                p = EditTextPreference(context).apply {
                                    when (any.getStr("summaryProvider")) {
                                        null -> summaryProvider = androidx.preference.EditTextPreference.SimpleSummaryProvider.getInstance()
                                        "PasswordSummaryProvider" -> summaryProvider = ProfileSettingsActivity.PasswordSummaryProvider
                                    }
                                    when (any.getStr("EditTextPreferenceModifiers")) {
                                        "Monospace" -> onBindEditTextListener = EditTextPreferenceModifiers.Monospace
                                        "Hosts" -> onBindEditTextListener = EditTextPreferenceModifiers.Hosts
                                        "Port" -> onBindEditTextListener = EditTextPreferenceModifiers.Port
                                        "Number" -> onBindEditTextListener = EditTextPreferenceModifiers.Number
                                    }
                                }
                            }
                            "SwitchPreference" -> {
                                p = SwitchPreference(context)
                            }
                            "SimpleMenuPreference" -> {
                                p = SimpleMenuPreference(context).apply {
                                    summaryProvider = androidx.preference.ListPreference.SimpleSummaryProvider.getInstance()
                                    val entries = any.optJSONObject("entries")
                                    if (entries != null) setMenu(this, entries)
                                }
                            }
                        }
                        // Set key & title
                        p.key = any.getString("key")
                        any.getStr("title")?.apply { p.title = this }
                        // Set icon
                        any.getStr("icon")?.apply {
                            p.icon = context.getDrawableByName(this)
                        }
                        // Set summary
                        any.getStr("summary")?.apply {
                            p.summary = this
                        }
                        // Add to category
                        preferenceCategory.addPreference(p)
                    }
                }
            }
        }

    fun setMenu(p: SimpleMenuPreference, entries: JSONObject) {
        val menuEntries = mutableListOf<String>()
        val menuEntryValues = mutableListOf<String>()
        entries.forEach { s, b ->
            menuEntryValues.add(s)
            menuEntries.add(b as String)
        }
        entries.apply {
            p.setEntries(menuEntries.toTypedArray())
            p.setEntryValues(menuEntryValues.toTypedArray())
        }
    }
}