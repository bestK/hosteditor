<script setup lang="ts">

import { onMounted, ref } from "vue";
import { Alert, Confirm, RequestUrl } from "wailsjs/go/main/App";
import { EditHosts, GetHostsBackupFilePaths, ReadHosts } from "wailsjs/go/services/hostsService";

import Codemirror from "codemirror-editor-vue3";
// placeholder
import "codemirror/addon/display/placeholder.js";
// language
import "codemirror/mode/javascript/javascript.js";
// placeholder
import "codemirror/addon/display/placeholder.js";
// theme
import "codemirror/theme/dracula.css";

import { useI18n } from "vue-i18n";



const { t, availableLocales: languages, locale } = useI18n();

const onclickLanguageHandle = (item: string) => {
  item !== locale.value ? (locale.value = item) : false;
};



const hostsValue = ref("");
const remoteHostUrl = ref("")
const isTypeing = ref(false)

const cmOptions = {
  mode: "text/plain",
  theme: "dracula",
};


onMounted(() => {
  ReadHosts().then((hosts) => {
    hostsValue.value = hosts;
  });
});


const onclickTypeRemoteHosts = () => {
  isTypeing.value = true
};

const onclickSave = async () => {
  const confirm = await Confirm(t("alert.edit"), t("alert.confirm-edit"))

  if (confirm === 'Yes' || confirm === 'Ok') {
    const res = await EditHosts(hostsValue.value)

    await Alert("提示", res)

    if (res !== 'Ok') {
      ReadHosts().then((hosts) => {
        hostsValue.value = hosts;
      });
    }
  }
};

const onclickRefresh = () => {
  ReadHosts().then(async (hosts) => {
    hostsValue.value = hosts;
    await Alert(t("alert.info"), t("alert.refreshed"))
  });
}

const onclickLoadRemoteHosts = async () => {
  const urlPattern = /^(https?|http):\/\/[^\s/$.?#].[^\s]*$/i;
  const isValid = urlPattern.test(remoteHostUrl.value);
  if (!isValid) {
    await Alert(t("alert.error"), `${t("alert.illegal-url")} ${remoteHostUrl.value}`)
  }

  const hosts = await RequestUrl(remoteHostUrl.value)

  if (!hosts || hosts == "") {
    return
  }

  const confirm = await Confirm(t("alert.edit"), t("alert.confirm-remote-hosts"))

  if (confirm === 'Yes' || confirm === 'Ok') {
    const res = await EditHosts(hosts)
    await Alert(t("alert.info"), res)
  }

  isTypeing.value = false
};

const onclickShowHostsBackup = async () => {
  const hostsPaths: string[] = await GetHostsBackupFilePaths()
  await Alert(t("alert.info"), hostsPaths.join("\n\r"))
}

</script>

<template>
  <!-- Header -->
  <div class="header">
    <!-- navigation -->
    <div class="nav">
      <router-link to="/">{{ t("nav.home") }}</router-link>
      <router-link to="/about">{{ t("nav.about") }}</router-link>
    </div>
    <!-- Menu -->
    <div class="menu" v-if="!isTypeing">
      <div class="language">
        <div v-for="item in languages" :key="item" :class="{ active: item === locale }"
          @click="onclickLanguageHandle(item)" class="lang-item">
          {{ t("languages." + item) }}
        </div>
      </div>
      <div class="bar">
        <div class="bar-btn" @click="onclickTypeRemoteHosts">
          {{ t("topbar.remoteHosts") }}
        </div>
        <div class="bar-btn" @click="onclickSave">
          {{ t("topbar.save") }}
        </div>
        <div class="bar-btn" @click="onclickRefresh">
          {{ t("topbar.refresh") }}
        </div>
        <div class="bar-btn" @click="onclickShowHostsBackup">
          {{ t("topbar.showHostsBackup") }}
        </div>
      </div>
    </div>
    <div class="menu" v-else>
      <input class="remote-host-url" v-model="remoteHostUrl" placeholder="https://..." />
      <div class="bar">
        <div class="bar-btn" @click="onclickLoadRemoteHosts">
          {{ t("topbar.loadRemoteHosts") }}
        </div>
        <div class="bar-btn" @click="onclickRefresh">
          {{ t("topbar.refresh") }}
        </div>
      </div>
    </div>
  </div>
  <div class="home">
    <Codemirror v-model:value="hostsValue" :options="cmOptions" />
  </div>
</template>

<style lang="scss">
.header {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  align-items: center;
  justify-content: space-between;
  height: 50px;
  padding: 0 10px;
  background-color: #333;

  .nav {
    a {
      display: inline-block;
      min-width: 50px;
      height: 30px;
      line-height: 30px;
      padding: 0 5px;
      margin-right: 8px;
      background-color: #282a36;
      border-radius: 2px;
      text-align: center;
      text-decoration: none;
      color: #000000;
      font-size: 14px;
      white-space: nowrap;

      &:hover,
      &.router-link-exact-active {
        background-color: #282a36;
        color: #00a67d;
      }
    }
  }

  .menu {
    display: flex;
    flex-direction: row;
    flex-wrap: nowrap;
    align-items: center;
    justify-content: space-between;

    .language {
      margin-right: 20px;
      background-color: #c3c3c3;
      overflow: hidden;

      .lang-item {
        display: inline-block;
        min-width: 50px;
        height: 30px;
        line-height: 30px;
        padding: 0 5px;
        background-color: #282a36;
        text-align: center;
        text-decoration: none;
        color: #111;
        font-size: 14px;
        // border: 1px solid #000000;

        &:hover {

          background-color: #282a36;
          color: #00a67d;
          cursor: not-allowed;
        }

        &.active {
          background-color: #282a36;
          color: #00a67d;
          cursor: pointer;
        }
      }
    }

    .bar {
      display: flex;
      flex-direction: row;
      flex-wrap: nowrap;
      align-items: center;
      justify-content: flex-end;
      min-width: 150px;

      .bar-btn {
        display: inline-block;
        min-width: 80px;
        height: 30px;
        line-height: 30px;
        padding: 0 5px;
        margin-left: 8px;
        background-color: #282a36;
        border-radius: 2px;
        text-align: center;
        text-decoration: none;
        color: #000000;
        font-size: 14px;

        &:hover {
          // background-color: #333;
          color: #00a67d;
          cursor: pointer;
        }
      }
    }

    .remote-host-url {
      width: fit-content;
      background-color: transparent;
      border: 2px solid #282a36;
      width: 400px;

    }

    input:focus {
      outline: none;
      color: #00a67d;
    }

  }
}



.home {
  height: calc(100% - 50px);
}

/* 隐藏浏览器默认的滚动条 */
::-webkit-scrollbar {
  width: 1px;
}

/* 滚动条的轨道（背景） */
::-webkit-scrollbar-track {
  background: transparent;
}

/* 滚动条的滑块（拖动条） */
::-webkit-scrollbar-thumb {
  background: #666;
  border-radius: 1px;
}

/* 当滑块被悬停时 */
::-webkit-scrollbar-thumb:hover {
  background: #555;
}
</style>
