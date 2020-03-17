<template>
  <v-app>
    <v-card width="400px" class="mx-auto mt-5">
      <v-card-title>
        <h1 class="display-1">Sign in</h1>
      </v-card-title>

      <v-card-text>
        <v-form>
          <v-text-field v-model="id" prepend-icon="mdi-account" type="text" label="id" />
          <v-text-field
            v-model="password"
            v-bind:type="showPassword ? 'text' : 'password'"
            prepend-icon="mdi-lock"
            v-bind:append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
            @click:append="showPassword = !showPassword"
            label="password"
          />
        </v-form>

        <v-card-actions>
          <v-btn class="info text-none" @click="login">Login</v-btn>
        </v-card-actions>
      </v-card-text>
    </v-card>
  </v-app>
</template>

<script>
export default {
  name: "Signin",
  data: function() {
    return {
      showPassword: false,
      id: "",
      password: ""
    };
  },

  methods: {
    login: function() {
      console.log("submit");
      let params = {
        id: this.id,
        password: this.password
      };
      this.$axios
        .post("/api/authentication", params)
        .then(res => {
          console.log(res);
        })
        .catch(err => {
          console.log(err);
        });
    }
  }
};
</script>