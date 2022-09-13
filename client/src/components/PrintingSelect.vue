<template>
  <div>
    <div class="source-choice" v-if="!usbInsert">
      <h1 class="display-1 mt-5 mb-5 source-choice-header text-center">
        Вставьте USB
      </h1>
    </div>
    <div class="source-choice" v-if="usbInsert && !usbReady">
      <h1 class="display-1 mt-5 mb-5 source-choice-header text-center">
        <div class="mb-3">Чтение носителя</div>
        <progress-bar
            size="huge"
            bar-color="#28a745"
            :bar-border-radius=5
            :val="processing"
        ></progress-bar>
      </h1>
    </div>
    <flash-Browser v-if="usbInsert && usbReady"></flash-Browser>
  </div>
</template>

<script>
import flashBrowser from "@/components/FlashBrowser";
import ProgressBar from 'vue-simple-progress'

export default {
  components: {
    flashBrowser,
    ProgressBar
  },
  data() {
    return {
      socket: null,
      usbInsert: false,
      usbReady: false,
      processing: 0
    }
  },
  methods: {
    wsMessage(message) {
      console.log(message)
      if (message === 'usb_insert') {
        this.usbInsert = true
      }
      if (message === 'usb_ready') {
        this.usbReady = true;
      }
      if (message === 'usb_conclusion') {
        this.usbInsert = false
        this.usbReady = false
        this.processing = 0
      }
      if (message.includes('processing:')) {
        this.processing = parseInt(message.replace('processing:', ''))
      }
    }
  },
  mounted() {
    let context = this;
    context.socket = new WebSocket("ws://localhost:8082/ws");
    context.socket.onmessage = function(event) {
      context.wsMessage(event.data)
    }
  }
}
</script>
<style scoped>
.source-choice {
  padding-top: 35vh;
}
.source-choice-header {
  font-family: 'Pacifico', cursive;
  text-shadow: 3px 3px 0px rgb(40, 167, 69, 1);
}
</style>