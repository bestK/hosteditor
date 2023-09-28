import { defineStore } from "pinia";
import { computed, ref } from "vue";

export const useHostsStore = defineStore("hosts", () => {
  const hosts = ref("");
  const getVlaue = computed(() => hosts.value);
  function setValue(value: string) {
    hosts.value = value;
  }

  return { hosts, getVlaue, setValue };
});


export const useHasNewRemoteHostsStore = defineStore("newRmoteHosts", () => {
  const hasNewHosts = ref(false);
  const getVlaue = computed(() => hasNewHosts.value);
  function setValue(value: boolean) {
    hasNewHosts.value = value;
  }

  return { hasNewHosts, getVlaue, setValue };
});
