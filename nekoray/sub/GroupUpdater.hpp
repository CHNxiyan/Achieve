#pragma once

#include "db/Database.hpp"

namespace NekoGui_sub {
    class RawUpdater {
    public:
        void updateClash(const QString &str);

        void update(const QString &str);

        int gid_add_to = -1; // 导入到指定组 -1 为当前选中组

        QList<std::shared_ptr<NekoGui::ProxyEntity>> updated_order; // 新增的配置，按照导入时处理的先后排序
    };

    class GroupUpdater : public QObject {
        Q_OBJECT

    public:
        void AsyncUpdate(const QString &str, int _sub_gid = -1, const std::function<void()> &finish = nullptr);

        void Update(const QString &_str, int _sub_gid = -1, bool _not_sub_as_url = false);

    signals:

        void asyncUpdateCallback(int gid);
    };

    extern GroupUpdater *groupUpdater;
} // namespace NekoGui_sub

// 更新所有订阅 关闭分组窗口时 更新动作继续执行
void UI_update_all_groups(bool onlyAllowed = false);
