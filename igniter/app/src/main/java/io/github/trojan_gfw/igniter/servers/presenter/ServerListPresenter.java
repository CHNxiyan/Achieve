package io.github.trojan_gfw.igniter.servers.presenter;

import android.content.Context;
import android.net.Uri;
import android.text.TextUtils;

import androidx.annotation.WorkerThread;

import com.stealthcopter.networktools.ping.PingStats;

import java.lang.ref.WeakReference;
import java.math.BigDecimal;
import java.util.LinkedHashSet;
import java.util.List;
import java.util.Set;

import io.github.trojan_gfw.igniter.Globals;
import io.github.trojan_gfw.igniter.TrojanConfig;
import io.github.trojan_gfw.igniter.TrojanURLHelper;
import io.github.trojan_gfw.igniter.TrojanURLParseResult;
import io.github.trojan_gfw.igniter.common.os.Task;
import io.github.trojan_gfw.igniter.common.os.Threads;
import io.github.trojan_gfw.igniter.common.sp.CommonSP;
import io.github.trojan_gfw.igniter.servers.contract.ServerListContract;
import io.github.trojan_gfw.igniter.servers.data.ServerListDataManager;
import io.github.trojan_gfw.igniter.servers.data.ServerListDataSource;

public class ServerListPresenter implements ServerListContract.Presenter {
    private final static String TAG = "ServerListPresenter";

    private final ServerListContract.View mView;
    private final ServerListDataSource mDataManager;
    private Set<TrojanConfig> mBatchDeleteConfigSet;

    public ServerListPresenter(ServerListContract.View view, ServerListDataSource dataManager) {
        mView = view;
        mDataManager = dataManager;
        view.setPresenter(this);
    }

    @Override
    public void hideImportFileDescription() {
        mView.dismissImportFileDescription();
    }

    @Override
    public void displayImportFileDescription() {
        mView.showImportFileDescription();
    }

    @Override
    public void exportServerListToFile() {
        Threads.instance().runOnWorkThread(new Task() {
            @Override
            public void onRun() {
                exportConfigFile();
            }
        });
    }

    private void exportConfigFile() {
        if (mDataManager.exportServers(Globals.getIgniterExportPath())) {
            mView.showExportServerListSuccess();
        } else {
            mView.showExportServerListFailure();
        }
    }

    @Override
    public void importConfigFromFile() {
        mView.openFileChooser();
    }

    @Override
    public void parseConfigsInFileStream(final Context context, final Uri fileUri) {
        mView.showLoading();
        Threads.instance().runOnWorkThread(new Task() {
            @Override
            public void onRun() {
                List<TrojanConfig> allConfigList = mDataManager.importServersFromFile(context, fileUri);
                mView.showServerConfigList(allConfigList);
                mView.showAddTrojanConfigSuccess();
                Threads.instance().runOnUiThread(mView::dismissLoading);
            }
        });
    }

    @Override
    public void batchOperateServerList() {
        if (mBatchDeleteConfigSet == null) {
            mBatchDeleteConfigSet = new LinkedHashSet<>();
        }
        mView.showServerListBatchOperation();
    }

    @Override
    public void selectServer(TrojanConfig config, boolean checked) {
        if (checked) {
            mBatchDeleteConfigSet.add(config);
        } else {
            mBatchDeleteConfigSet.remove(config);
        }
    }

    @Override
    public void selectAll(List<TrojanConfig> configList) {
        mBatchDeleteConfigSet.addAll(configList);
        mView.selectAllServers();
    }

    @Override
    public void exitServerListBatchOperation() {
        mView.hideServerListBatchOperation();
        mBatchDeleteConfigSet.clear();
    }

    @Override
    public void deselectAll(List<TrojanConfig> configList) {
        mBatchDeleteConfigSet.clear();
        mView.deselectAllServers();
    }

    @Override
    public void saveServerList(List<TrojanConfig> configList) {
        mView.showLoading();
        Threads.instance().runOnWorkThread(new Task() {
            @Override
            public void onRun() {
                mDataManager.replaceServerConfigs(configList);
                Threads.instance().runOnUiThread(mView::dismissLoading);
            }
        });
    }

    @Override
    public void batchDelete() {
        if (mBatchDeleteConfigSet.isEmpty()) {
            mView.showBatchDeletionSuccess();
        } else {
            // only remove items in view here.
            // Actual deletion should be done after exiting batch operation mode.
            mView.batchDelete(mBatchDeleteConfigSet);
            mBatchDeleteConfigSet.clear();
            mView.showBatchDeletionSuccess();
        }
    }

    @Override
    public void addServerConfig(final String trojanUrl) {
        Threads.instance().runOnWorkThread(new Task() {
            @Override
            public void onRun() {
                TrojanURLParseResult parseResult = TrojanURLHelper.ParseTrojanURL(trojanUrl);
                if (parseResult != null) {
                    TrojanConfig newConfig = TrojanURLHelper.CombineTrojanURLParseResultToTrojanConfig(parseResult, Globals.getTrojanConfigInstance());
                    mDataManager.saveServerConfig(newConfig);
                    loadConfigs();
                    mView.showAddTrojanConfigSuccess();
                } else {
                    mView.showQRCodeScanError(trojanUrl);
                }
            }
        });
    }

    @Override
    public void handleServerSelection(TrojanConfig config) {
        mView.selectServerConfig(config);
    }

    @Override
    public void displaySubscribeSettings() {
        mView.showSubscribeSettings(CommonSP.getServerSubscribeUrl(""));
    }

    @Override
    public void updateSubscribeServers() {
        String url = CommonSP.getServerSubscribeUrl("");
        if (TextUtils.isEmpty(url)) {
            mView.showSubscribeUpdateFailed();
            return;
        }
        mView.showLoading();
        Threads.instance().runOnWorkThread(new Task() {
            @Override
            public void onRun() {
                mDataManager.requestSubscribeServerConfigs(url, new SubscribeCallback(ServerListPresenter.this));
            }
        });
    }

    @Override
    public void saveSubscribeSettings(String url) {
        CommonSP.setServerSubscribeUrl(url);
    }

    @Override
    public void hideSubscribeSettings() {
        mView.dismissSubscribeSettings();
    }

    @Override
    public void start() {
        Threads.instance().runOnWorkThread(new Task() {
            @Override
            public void onRun() {
                loadConfigs();
            }
        });
    }

    @Override
    public void gotoScanQRCode(boolean fromGallery) {
        if (fromGallery) {
            mView.scanQRCodeFromGallery();
        } else {
            mView.scanQRCodeFromCamera();
        }
    }

    @Override
    public void pingAllProxyServer(List<TrojanConfig> configList) {
        for (TrojanConfig config : configList) {
            mDataManager.pingTrojanConfigServer(config, new PingServerCallBack(this));
        }
    }

    @WorkerThread
    private void loadConfigs() {
        List<TrojanConfig> trojanConfigs = mDataManager.loadServerConfigList();
        mView.showServerConfigList(trojanConfigs);
    }

    @WorkerThread
    private void showSubscribeServersSuccess() {
        Threads.instance().runOnUiThread(() -> {
            mView.dismissLoading();
            mView.showSubscribeUpdateSuccess();
        });
        loadConfigs();
    }

    private void showSubscribeServersFailed() {
        Threads.instance().runOnUiThread(() -> {
            mView.dismissLoading();
            mView.showSubscribeUpdateFailed();
        });
    }

    private void setPingDelayTime(TrojanConfig config, float pingDelayTime) {
        Threads.instance().runOnUiThread(() -> {
            mView.setPingServerDelayTime(config, pingDelayTime);
        });
    }

    private static class SubscribeCallback implements ServerListDataSource.Callback {
        private final WeakReference<ServerListPresenter> mPresenterRef;

        public SubscribeCallback(ServerListPresenter presenter) {
            mPresenterRef = new WeakReference<>(presenter);
        }

        @Override
        public void onSuccess() {
            ServerListPresenter presenter = mPresenterRef.get();
            if (presenter != null) {
                presenter.showSubscribeServersSuccess();
            }
        }

        @Override
        public void onFailed() {
            ServerListPresenter presenter = mPresenterRef.get();
            if (presenter != null) {
                presenter.showSubscribeServersFailed();
            }
        }
    }

    private static class PingServerCallBack implements ServerListDataSource.PingCallback {
        private final WeakReference<ServerListPresenter> mPresenterRef;

        public PingServerCallBack(ServerListPresenter presenter) {
            mPresenterRef = new WeakReference<>(presenter);
        }

        @Override
        public void onSuccess(TrojanConfig config, PingStats pingStats) {
            ServerListPresenter presenter = mPresenterRef.get();
            if (presenter != null) {
                //LogHelper.d(TAG, "ping "+ config.getRemoteAddr() + ": "+ pingStats.toString());
                if (!pingStats.isReachable()) {
                    presenter.setPingDelayTime(config, ServerListDataManager.SERVER_UNABLE_TO_REACH);
                    return;
                }
                BigDecimal b = BigDecimal.valueOf(pingStats.getAverageTimeTaken());
                float pingDelayTime = b.setScale(1, BigDecimal.ROUND_HALF_UP).floatValue();
                presenter.setPingDelayTime(config, pingDelayTime);
            }
        }

        @Override
        public void onFailed(TrojanConfig config) {
            ServerListPresenter presenter = mPresenterRef.get();
            if (presenter != null) {
                presenter.setPingDelayTime(config, ServerListDataManager.SERVER_UNABLE_TO_REACH);
            }
        }
    }
}
