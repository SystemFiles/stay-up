<template>
  <div class="home-page" :key="componentKey">
    <div class="banner">
      <div class="container" style="margin-top: 2.4rem;">
        <h1 class="d-flex justify-content-center logo-font">StayUp</h1>
        <p
          class="d-flex justify-content-center"
        >Ultra-Lightweight Service / Host Monitoring Solution</p>
      </div>
    </div>
    <div class="container page">
      <h2 class="mt-5">Services</h2>
      <p>a list of services being tracked</p>

      <div class="container mt-5 mx-auto">
        <div class="row">
          <div class="col-md-3" v-for="svc in services">
            <service-list-item
              v-bind:id="svc.id"
              v-bind:name="svc.name"
              v-bind:description="svc.description"
              v-bind:host="svc.host"
              v-bind:port="svc.port"
              v-bind:currentStatus="svc.currentStatus"
              @deleteSvc="deleteSvc(svc)"
              @updateSvc="updateSvc(svc)"
            />
          </div>
        </div>
      </div>
    </div>
    <div class="container operations fixed-bottom">
      <div class="row bg-body pb-1 justify-content-end">
        <div class="col-md-2 col-sm-12">
          <router-link :to="{ name: 'add' }">
            <button type="button" class="w-100 my-1 btn btn-success">ADD</button>
          </router-link>
        </div>
      </div>
    </div>
    <div class="row">
      <div class="fixed-bottom mb-5 ml-5 d-flex justify-content-end">
        <div v-if="error" class="toast show" role="alert" aria-live="assertive" aria-atomic="true">
          <div class="toast-header">
            <strong class="me-auto">Error</strong>
            <button
              type="button"
              class="btn-close ms-2 mb-1"
              v-on:click="toggleError"
              data-bs-dismiss="toast"
              aria-label="Close"
            >
              <span aria-hidden="true"></span>
            </button>
          </div>
          <div class="toast-body">{{ errorMessage }}</div>
        </div>
      </div>
    </div>
    <div class="row">
      <div class="fixed-bottom mb-5 ml-5 d-flex justify-content-end">
        <div
          v-if="success"
          class="toast show"
          role="alert"
          aria-live="assertive"
          aria-atomic="true"
        >
          <div class="toast-header">
            <strong class="me-auto">StayUp</strong>
            <button
              type="button"
              class="btn-close ms-2 mb-1"
              v-on:click="toggleSuccess"
              data-bs-dismiss="toast"
              aria-label="Close"
            >
              <span aria-hidden="true"></span>
            </button>
          </div>
          <div class="toast-body">{{ successMessage }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import ServiceListItemVue from "../components/ServiceListItem.vue";
import { SvcService } from "../common/api.service";
export default {
  name: "Home",
  data() {
    return {
      services: [{
        id: 1,
        name: "Bitwarden",
        host: "bitwarden.sykesdev.ca",
        port: 443,
        description: "A self-host password vault",
        currentStatus: "SLOW"
      }, {
        id: 2,
        name: "Jellyfin",
        host: "jely.sykeshome.io",
        port: 443,
        description: "Media server for TV/Movies",
        currentStatus: "UP"
      }],
      error: false,
      errorMessage: '',
      success: false,
      successMessage: '',
      componentKey: 0
    }
  },
  components: {
    serviceListItem: ServiceListItemVue
  },
  methods: {
    toggleError() {
      this.$data.error = false
    },
    toggleSuccess() {
      this.$data.success = false
    },
    refresh() {
      this.componentKey += 1 // will force re-render
    },
    getServices() {
      // Open websocket and receive data to services state
    },
    async updateSvc(svc) {
      console.log("NOT IMPLEMENTED")
      this.$data.error = true
      this.$data.errorMessage = "Feature not implemented yet..."
    },
    async deleteSvc(svc) {
      SvcService.delete(svc.id).then((res) => {
        if (res.status === 200) {
          this.$data.success = true
          this.$data.successMessage = `Successfully deleted service with ID, ${svc.id}!`
        }
      }).catch((err) => {
        this.$data.error = true
        this.$data.errorMessage = `Failed to delete service with ID, ${svc.id}. See log for details.`
        console.log(err.response ? err.response.data.message : err.message)
      })
    }
  }
};
</script>
