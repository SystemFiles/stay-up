<template>
  <div class="home-page" :key="componentKey">
    <div class="banner">
      <div class="container" style="margin-top: 2.4rem;">
        <h1 class="d-flex justify-content-center logo-font">StayUp</h1>
        <p class="d-flex justify-content-center ">Ultra-Lightweight Service / Host Monitoring Solution</p>
      </div>
    </div>
    <div class="container page">
      <h2 class="mt-5">Services</h2>
      <p>a list of services being tracked</p>

      <div class="container mt-5 mx-auto">
        <div class="row">
          <div class="col-md-3">
            <div class="card my-2">
              <div class="card-body">
                <h4 class="card-title">Jellyfin</h4>
                <h6 style="font-size: 0.6rem" class="card-subtitle mb-2 text-muted">jelly.sykeshome.io : 443</h6>
                <p class="card-text">No description was provided</p>
                <div class="row mx-0">
                  <div class="badge bg-success px-5 py-3 mx-auto">UP</div>
                </div>
                <div class="row mt-3">
                  <div class="col-md-6">
                    <button type="button" class="w-100 my-1 btn btn-danger">DELETE</button>
                  </div>
                  <div class="col-md-6">
                    <button type="button" class="w-100 my-1 btn btn-primary">EDIT</button>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="col-md-3">
            <div class="card my-2">
              <div class="card-body">
                <h4 class="card-title">Deluge</h4>
                <h6 style="font-size: 0.6rem" class="card-subtitle mb-2 text-muted">deluge.sykeshome.io : 3000</h6>
                <p class="card-text">No description was provided</p>
                <div class="row mx-0">
                  <div class="badge bg-danger px-5 py-3 mx-auto">DOWN</div>
                </div>
                <div class="row mt-3">
                  <div class="col-md-6">
                    <button type="button" class="w-100 my-1 btn btn-danger">DELETE</button>
                  </div>
                  <div class="col-md-6">
                    <button type="button" class="w-100 my-1 btn btn-primary">EDIT</button>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="col-md-3">
            <div class="card my-2">
              <div class="card-body">
                <h4 class="card-title">Resume Site</h4>
                <h6 style="font-size: 0.6rem" class="card-subtitle mb-2 text-muted">sykesdev.ca : 443</h6>
                <p class="card-text">No description was provided</p>
                <div class="row mx-0">
                  <div class="badge bg-warning px-5 py-3 mx-auto">SLOW</div>
                </div>
                <div class="row mt-3">
                  <div class="col-md-6">
                    <button type="button" class="w-100 my-1 btn btn-danger">DELETE</button>
                  </div>
                  <div class="col-md-6">
                    <button type="button" class="w-100 my-1 btn btn-primary">EDIT</button>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="col-md-3">
            <div class="card my-2">
              <div class="card-body">
                <h4 class="card-title">Bitwarden</h4>
                <h6 style="font-size: 0.6rem" class="card-subtitle mb-2 text-muted">bitwarden.sykesdev.ca : 443</h6>
                <p class="card-text">No description was provided</p>
                <div class="row mx-0">
                  <div class="badge bg-success px-5 py-3 mx-auto">UP</div>
                </div>
                <div class="row mt-3">
                  <div class="col-md-6">
                    <button type="button" class="w-100 my-1 btn btn-danger">DELETE</button>
                  </div>
                  <div class="col-md-6">
                    <button type="button" class="w-100 my-1 btn btn-primary">EDIT</button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="container operations fixed-bottom">
      <div class="row bg-body pb-1">
        <div class="col-md-6 col-sm-12">
          <button type="button" v-on:click="refresh" class="w-100 my-1 btn btn-primary">Refresh Services</button>
        </div>
        <div class="col-md-6 col-sm-12">
          <router-link :to="{ name: 'add' }">
            <button type="button" class="w-100 my-1 btn btn-success">ADD Service</button>
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Home",
  data() {
    return {
      componentKey: 0
    }
  },
  methods: {
    refresh() {
      this.componentKey += 1
      console.log(this.componentKey)
    }
  }
};
</script>
