<template>
  <div class="add-page">
    <div class="banner">
      <div class="container" style="margin-top: 2.4rem;">
        <h1 class="d-flex justify-content-center logo-font">StayUp</h1>
        <p class="d-flex justify-content-center ">Ultra-Lightweight Service / Host Monitoring Solution</p>
      </div>
    </div>
    <div class="container page">
      <bar-loader class="w-100 position-fixed fixed-bottom" :loading="isLoading" :size="150"></bar-loader>
      <div class="row mt-5">
        <form data-bitwarden-watching="1" @submit.prevent="onPublish">
          <fieldset>
            <legend>Add a Service</legend>
            <div class="form-group">
              <label for="serviceName" class="form-label mt-4">Service Name</label>
              <input type="name" class="form-control" v-model="serviceName" id="serviceNameInput" aria-describedby="nameHelp" placeholder="enter a name for the service">
              <small id="nameHelp" class="form-text text-muted">Try to keep the name as short as possible</small><br/>
              <label for="serviceDesc" class="form-label mt-4">Short Description</label>
              <input type="text" class="form-control" v-model="serviceDesc" id="serviceDescInput" placeholder="enter a short description of the service">
            </div>
            <div class="form-group">
              <label for="serviceHost" class="form-label mt-4">Host and Port</label>
              <div class="row">
                <div class="col-sm-8 col-8">
                  <input type="text" class="form-control" v-model="serviceHost" id="serviceHostInput" placeholder="IPv4 or Hostname">
                </div>
                <div class="col-sm-4 col-4">
                  <input type="number" class="form-control" v-model="servicePort" id="servicePortInput" placeholder=": port">
                </div>
              </div>
            </div>
            <div class="form-group">
              <label for="serviceProtocol" class="form-label mt-4">Communication Protocol</label>
              <select class="form-select" v-model="serviceProtocol" id="serviceProtocol">
                <option value="TCP">Transmission Control Protocol (TCP)</option>
                <option value="UDP">User Datagram Protocol (UDP)</option>
              </select>
            </div>
            <div class="form-group mt-4">
              <fieldset class="form-group">
                  <label for="serviceTimeout" class="form-label">Timeout (ms)</label>
                  <input type="range" class="form-range" v-model="serviceTimeout" min="0" max="1000" step="50" id="serviceTimeout">
                  <p class="d-flex justify-content-end">{{serviceTimeout}} ms</p>
              </fieldset>
            </div>
            <br/>
            <button type="submit" class="btn btn-primary">Submit</button>
            <router-link :to="{ name: 'home' }" replace>
              <button type="button" class="btn btn-secondary">Cancel</button>
            </router-link>
          </fieldset>
        </form>
      </div>
    </div>
    <div class="row">
      <div class="fixed-bottom mb-5 ml-5 d-flex justify-content-end">
        <div v-if="error" class="toast show" role="alert" aria-live="assertive" aria-atomic="true">
          <div class="toast-header">
            <strong class="me-auto">Error</strong>
            <button type="button" class="btn-close ms-2 mb-1" v-on:click="toggleError" data-bs-dismiss="toast" aria-label="Close">
              <span aria-hidden="true"></span>
            </button>
          </div>
          <div class="toast-body">
            {{errorMessage}}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { BarLoader } from "@saeris/vue-spinners"
import Vue from "vue"
import store from "@/store"
import axios from "axios"
import { mapGetters } from "vuex"
import { ADDSVC } from "@/store/mutations.type"
import { SvcService } from "@/common/api.service"

export default {
  name: "Add",
  data() {
    return {
      isLoading: false,
      error: false,
      errorMessage: '',
      serviceName: '',
      serviceHost: '',
      servicePort: '',
      serviceDesc: '',
      serviceProtocol: '',
      serviceTimeout: 250
    }
  },
  components: {
    BarLoader
  },
  methods: {
    toggleError() {
      this.$data.error = false
    },
    async onPublish() {
      const reqData = {
        name: this.$data.serviceName.trim().substring(0,25),
        description: this.$data.serviceDesc.length > 0 ? this.$data.serviceDesc.trim().substring(0,120) : "No description was provided",
        host: this.$data.serviceHost.trim(),
        port: parseInt(this.$data.servicePort),
        protocol: this.$data.serviceProtocol.trim(),
        timeout: parseInt(this.$data.serviceTimeout)
      }

      // start loading
      this.$data.isLoading = true

      SvcService.post(reqData).then((res) => {
        this.$data.isLoading = false
        this.$router.replace({name: 'home'})
      }).catch((err) => {
        console.log(err.response ? err.response.data.message : err.message)
        this.$data.isLoading = false
        this.$data.error = true
        this.$data.errorMessage = `Could not create new service. See log for details.`
      })

      // await this.axios.post("http://localhost:5555/api/service", reqData).then((res) => {
      //   this.$data.isLoading = false
      //   this.$router.replace({name: 'home'})
      // }).catch((err) => {
      //   console.log(err.response ? err.response.data.message : err.message)
      //   this.$data.isLoading = false
      //   this.$data.error = true
      //   this.$data.errorMessage = `Could not create new service. See log for details.`
      // })
    }
  }
};
</script>
