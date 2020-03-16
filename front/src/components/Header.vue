<template>
  <div class="home">
    <v-app-bar color="blue lighten-2">
      <v-toolbar-title>
        <v-btn class="ml-4 text-none" outlined to="/">Top</v-btn>
      </v-toolbar-title>
      <v-btn class="ml-4 text-none" outlined to="/about">About</v-btn>
      <v-spacer></v-spacer>
      <v-btn class="mr-4 text-none" outlined to="signin">Sign in</v-btn>
      <v-btn class="text-none" to="signup">Sign up</v-btn>
    </v-app-bar>
  </div>
</template>

<script>
export default {};
</script>
