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
                <v-text-field v-model="url" prepend-icon="mdi-link-variant" type="url" label="URL" />
              </v-form>

              <v-card-actions>
                <v-btn class="info text-none" @click="send">Send</v-btn>
              </v-card-actions>
            </v-card-text>
          </v-card>
        </v-tab-item>
      </v-tabs>
    </v-container>
  </v-app>
</template>

<script>
export default {
  name: "Dashboard",
  data: function() {
    return {
      url: ""
    };
  },

  methods: {
    send: function() {
      console.log("send");
      this.$axios
        .post("/api/hogehoge", this.url)
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

<style scoped>
.container {
  max-width: 900px;
}
</style>
