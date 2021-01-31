<template>
  <div>
    <div class="media mt-1">
      <h3 class="h3-responsive font-weight-bold mr-3 pt-0">{{ name }}</h3>
      <div class="media-body mb-3 mb-lg-3">
        <mdb-badge
          @click.native="onDelete"
          tag="a"
          color="danger-color"
          class="ml-2 float-right"
          >-</mdb-badge
        >
        <h6 v-if="tags" class="mt-0 font-weight-bold">
          <a v-for="(tag, idx) in tags" :key="idx" :href="'/tags/' + tag.uid">{{
            tag.name
          }}</a>
        </h6>
        <hr class="hr-bold my-2" />
      </div>
    </div>
  </div>
</template>

<script>
import { mdbBadge } from "mdbvue";

export default {
  name: "Recipe",
  components: {
    mdbBadge
  },
  props: {
    index: {
      type: Number
    },
    uid: {
      type: String
    },
    name: {
      type: String
    },
    tags: {
      type: Array
    }
  },
  methods: {
    onDelete() {
      this.$emit("delete", this.index);
    }
  }
};
</script>

<style scoped></style>
