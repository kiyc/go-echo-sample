<template>
  <v-container>
    <v-row class="mt-6">
      <v-col sm="6" offset-sm="3">
        <form method="post" action="/login" id="login-form">
          <v-card>
            <v-card-title>Go Echo Sample</v-card-title>
            <v-card-text>
              <v-text-field type="text" v-model="account" label="Account" @keyup.enter="login"></v-text-field>
              <v-text-field
                type="password"
                v-model="password"
                label="Password"
                @keyup.enter="login"
              ></v-text-field>
              <h3 class="red--text" v-if="hasErrors">Login failed.</h3>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn type="button" @click="login">Login</v-btn>
            </v-card-actions>
          </v-card>
        </form>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from "axios";
import { mapMutations } from "vuex";

export default {
  data() {
    return {
      account: null,
      password: null,
      hasErrors: false
    };
  },
  methods: {
    ...mapMutations(["setAuthToken"]),
    login() {
      this.hasErrors = false;
      const data = new FormData();
      data.append("account", this.account);
      data.append("password", this.password);
      axios({
        method: "post",
        url: "/login",
        data
      })
        .then(response => {
          if (response.status === 200) {
            this.setAuthToken(response.data.token);
            this.$router.push("/dashboard");
          }
        })
        .catch(error => {
          this.hasErrors = true;
          console.log(error.response);
        });
    }
  }
};
</script>