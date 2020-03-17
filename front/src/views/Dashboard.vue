<template>
  <v-app>
    <v-container>
      <v-tabs fixed-tabs background-color="blue lighten-3">
        <v-tab class="text-none">My-article</v-tab>
        <v-tab-item>
          <h1>ユーザーの記事一覧</h1>
          <span>記事毎にeditボタンをつけよう</span>
        </v-tab-item>
        <v-tab class="text-none">New</v-tab>
        <v-tab-item>
          <v-card width="500px" class="mx-auto mt-5">
            <v-card-title>
              <h1 class="display-1">自分の記事を追加</h1>
            </v-card-title>

            <v-card-text>
              <v-form>
                <v-text-field v-model="email" prepend-icon="mdi-email" type="email" label="mail" />
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
        </v-tab-item>
      </v-tabs>
    </v-container>
  </v-app>
</template>

<style scoped>
.container {
  max-width: 900px;
}
</style>
