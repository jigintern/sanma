<template>
  <v-app>
    <v-container>
      <v-tabs fixed-tabs background-color="blue lighten-3">
        <v-tab class="text-none">My-article</v-tab>
        <v-tab-item>
          <div v-show="loading">Loading...</div>
          <v-row dense v-show="!loading">
            <v-col v-for="(article, index) in articles" :key="index" cols="12">
              <v-card
                color="blue lighten-3"
                height="100%"
                :href="article.article_url"
                target="_blank"
              >
                <v-card-title class="headline">{{ article.title }}</v-card-title>

                <v-card-text>id: {{ article.user.id }}</v-card-text>

                <v-card-actions>
                  <v-btn class="text-none" @click="edit">Edit</v-btn>
                </v-card-actions>
              </v-card>
            </v-col>
          </v-row>
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
      loading: true,
      url: "",
      articles: []
    };
  },

  mounted: function() {
    this.$axios.get("https://qiita.com/api/v2/items").then(res => {
      console.log(res);
      this.articles = res.data;
      this.loading = false;
    });
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
