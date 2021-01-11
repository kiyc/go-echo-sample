<template>
  <v-container>
    <v-row>
      <v-col cols="6">
        <h3>Go Echo Sample</h3>
      </v-col>
      <v-col cols="6">
        <p class="text-right mb-0">
          <a @click="logout">Logout</a>
        </p>
      </v-col>
    </v-row>
    <h1>Dashboard</h1>
    <p>
      <router-link to="/users">Users</router-link>
    </p>
  </v-container>
</template>

<script>
import { mapActions } from "vuex";

export default {
  methods: mapActions(["logout"])
};
</script>