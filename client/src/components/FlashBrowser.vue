<template>
  <div>
    <div class="source-choice">
      <h1 class="display-1 mt-5 mb-5 source-choice-header text-center" v-if="!loading && files !== null && files.length > 0">
        Выберите файл
      </h1>
      <h1 class="display-1 mt-5 mb-5 source-choice-header text-center" v-if="loading && files === null">
        Готовим список
      </h1>
      <h1 class="display-1 mt-5 mb-5 source-choice-header text-center" v-if="!loading && files !== null && files.length === 0">
        Файлы не найдены
      </h1>
      <div class="file-support-legend" v-if="!loading && files !== null && files.length === 0">
        Для печати подходят файлы пакета Microsoft Office и совместимые<br>
        .doc, .docx, .xls, .xlsx, .ppt, .pptx
      </div>
    </div>
    <div class="file-list-window" v-if="!loading && files !== null && files.length > 0">
      <ul>
        <li v-for="item in computedFiles" :key="item.index" class="file-list-elem">
          <div class="row" v-if="item.size > 0">
            <div class="col-4 file-list-elem-left-side">{{ item.filename }}</div>
            <div class="col-4 file-list-elem-center-side">{{ formatBytes(item.size, 2) }}</div>
            <div class="col-4 file-list-elem-right-side">
              <div class="btn btn-success" @click="markAsPrinted(item.index, !item.printing)" v-if="!item.printing">Распечатать</div>
              <div class="btn btn-danger" @click="markAsPrinted(item.index, !item.printing)" v-if="item.printing">Не печатать</div>
            </div>
          </div>
        </li>
      </ul>
    </div>
    <div class="file-start-printing-process" v-if="hasPrintingItem">
      <div class="btn btn-light btn-lg">Продолжить</div>
    </div>
  </div>
</template>

<script>
import axios from "axios";



export default {
  name: "FlashBrowser",
  data() {
    return {
      loading: true,
      files: null
    }
  },
  computed: {
    computedFiles() {
      if (this.files !== null) {
        return this.files.map((_, index) => ({
          ..._,
          index: index,
        }))
      }
      return null
    },
    hasPrintingItem() {
      let has = false
      if (this.files !== null) {
        this.files.forEach(function (item) {
          if (item.printing) {
            has = true
          }
        })
      }
      return has
    }
  },
  methods: {
    markAsPrinted(index, printed) {
      this.files[index].printing = printed
    },
    prepareData(date) {
      this.files = date.map((_) => ({
        ..._,
        printing: false,
      }))
    },
    formatBytes(bytes, decimals = 2) {
      if (bytes === 0) return '0 Bytes';
      let k = 1024;
      let dm = decimals < 0 ? 0 : decimals;
      let sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];
      let i = Math.floor(Math.log(bytes) / Math.log(k));
      return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + ' ' + sizes[i];
    },
    sleep(milliseconds) {
      const date = Date.now();
      let currentDate = null;
      do {
        currentDate = Date.now();
      } while (currentDate - date < milliseconds);
    },
  },
  mounted() {
    this.loading = true
    axios.get('http://localhost:8081/usb-list').then((response) => {
      this.prepareData(response.data)
      this.sleep(1200)
      this.loading = false
    })
  }
}
</script>

<style scoped>
.file-support-legend {
  font-size: 24px;
  text-align: center;
  width: 470px;
  padding: 21px 0px;
  margin: 0 auto;
}
.file-start-printing-process {
  margin-top: 15px;
  text-align: center;
}
.file-list-elem-center-side {
  line-height: 45px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
.file-list-elem-right-side {
  line-height: 45px;
  text-align: right;
  padding-right: 5px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
.file-list-elem-left-side {
  font-size: 25px;
  line-height: 50px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
.file-list-elem {
  list-style: none;
}
.file-list-elem .row {
  color: #28a745;
  border: 1px solid #fff;
  width: 98%;
  margin-bottom: 5px !important;
  background: #fff;
  border-radius: 3px;
  margin: 0 auto;
  height: 50px !important;
}
.file-list-window ul {
  padding: 0px !important;
  margin: 0px !important;
}
.file-list-window {
  background: rgba(255,255,255,.8);
  border-radius: 5px;
  padding: 10px 0px 5px 0px;
  box-shadow: 0px 3px 0px rgba(40, 167, 69, 0.5);
  border: 1px solid;
}
.source-choice {
  padding-top: 10vh;
}
.source-choice-header {
  font-family: 'Pacifico', cursive;
  text-shadow: 3px 3px 0px rgb(40, 167, 69, 1);
}
</style>