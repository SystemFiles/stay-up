<template>
  <div class="add-page">
    <div class="banner">
      <div class="container" style="margin-top: 2.4rem;">
        <h1 class="d-flex justify-content-center logo-font">StayUp</h1>
        <p class="d-flex justify-content-center ">Ultra-Lightweight Service / Host Monitoring Solution</p>
      </div>
    </div>
    <div class="container page">
      <div class="row mt-5">
        <form data-bitwarden-watching="1">
          <fieldset>
            <legend>Add a Service</legend>
            <div class="form-group">
              <label for="serviceName" class="form-label mt-4">Service Name</label>
              <input type="name" class="form-control" v-model="serviceName" id="serviceNameInput" aria-describedby="nameHelp" placeholder="enter a name for the service">
              <small id="nameHelp" class="form-text text-muted">Try to keep the name as short as possible</small><br/>
              <label for="serviceDesc" class="form-label mt-4">Short Description <span class="text-muted">(optional)</span></label>
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
                <option>Transmission Control Protocol (TCP)</option>
                <option>User Datagram Protocol (UDP)</option>
              </select>
            </div>
            <div class="form-group mt-4">
              <fieldset class="form-group">
                  <label for="serviceTimeout" class="form-label">Timeout (ms)</label>
                  <input type="range" class="form-range" v-model="serviceTimeout" min="0" max="1000" step="50" id="serviceTimeout">
                  <p class="d-flex justify-content-end">50 ms</p>
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
  </div>
</template>

<script>
export default {
  name: "Add",
  data() {
    return {
      serviceName: null,
      serviceHost: null,
      servicePort: null,
      serviceDesc: null,
      serviceProtocol: null,
      serviceTimeout: null
    }
  },
  methods: {
    onSubmit(serviceName, serviceHost, serviceDesc, servicePort, serviceProtocol, serviceTimeout) {
      this.$store
        .dispatch(ADDSVC, {serviceName, serviceDesc, serviceHost, servicePort, serviceProtocol, serviceTimeout})
        .then(() => this.$router.replace({name: "home"}))
    }
  }
};
</script>
