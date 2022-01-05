<template>
  <div class="card my-2">
    <div class="card-body">
      <h4 class="card-title">{{ name.trim().substring(0, 17) }}</h4>
      <h6 style="font-size: 0.6rem" class="card-subtitle mb-2 text-muted">
        {{ host }}:{{ port }}
      </h6>
      <p class="card-text">
        {{ description.trim().substring(0, 28) + " ..." }}
      </p>
      <div class="row mx-0">
        <div
          v-bind:class="{
            'bg-success': currentStatus === 'UP',
            'bg-warning': currentStatus === 'SLOW',
            'bg-danger': currentStatus === 'DOWN'
          }"
          class="badge px-5 py-3 mx-auto"
        >
          {{ currentStatus }}
        </div>
      </div>
      <div class="row mt-3">
        <div class="col-md-6">
          <button
            @click="deleteSvc"
            type="button"
            class="w-100 my-1 btn btn-danger"
          >
            DELETE
          </button>
        </div>
        <div class="col-md-6">
          <button
            @click="updateSvc"
            type="button"
            class="w-100 my-1 btn btn-primary"
          >
            EDIT
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ServiceListItem",
  props: {
    id: {
      type: Number,
      required: true
    },
    name: {
      type: String,
      required: true
    },
    host: {
      type: String,
      required: true
    },
    port: {
      type: Number,
      required: true
    },
    description: {
      type: String,
      required: false,
      default: "No description was provided"
    },
    currentStatus: {
      type: String,
      required: true
    }
  },
  methods: {
    deleteSvc() {
      this.$emit("deleteSvc");
    },
    updateSvc() {
      this.$emit("updateSvc");
    }
  }
};
</script>
