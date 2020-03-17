<template>
  <v-app>
    <v-card width="400px" class="mx-auto mt-5">
      <v-card-title>
        <h1 class="display-1">Sign in</h1>
      </v-card-title>

      <v-card-text>
        <v-form>
          <v-text-field prepend-icon="mdi-email" type="email" label="mail" />
          <v-text-field prepend-icon="mdi-lock" type="password" label="password" />
        </v-form>

        <v-card-actions>
          <v-btn class="info text-none">Login</v-btn>
        </v-card-actions>
      </v-card-text>
    </v-card>
  </v-app>
</template>
