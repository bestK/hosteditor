import { createI18n } from "vue-i18n";

import en from "./locales/en.json";
import zhHans from "./locales/zh-Hans.json";

const i18n = createI18n({
  locale: "en",
  fallbackLocale: "en",
  legacy: false,
  messages: {
    "zh-Hans": zhHans,
    en: en,
  },
});

export default i18n;
