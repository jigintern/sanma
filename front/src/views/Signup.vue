<template>
  <v-app>
    <v-card width="400px" class="mx-auto mt-5">
      <v-card-title>
        <h1 class="display-1">Sign up</h1>
      </v-card-title>

      <v-card-text>
        <div class="display-1 font-weight-black">ID:{{ splitMail }}</div>
        <v-form>
          <v-text-field v-model="mail" prepend-icon="mdi-email" type="email" label="mail" />
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
      mail: "",
      password: ""
    };
  },

  computed: {
    splitMail: function() {
      let account = this.mail.split("@")[0];
      return account;
    }
  },

  methods: {
    login: function() {
      console.log("submit");
      let params = {
        id: this.splitMail,
        mail: this.mail,
        password: this.password
      };
      console.log(params);
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