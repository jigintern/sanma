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
          <h1>new</h1>
          <span>add crawl</span>
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
