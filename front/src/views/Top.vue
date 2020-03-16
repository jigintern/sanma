<template>
  <div class="home">
    <h1>Top page</h1>

    <v-container>
      <v-container>
        <v-row dense>
          <v-col
            v-for="(article, index) in articles"
            :key="index"
            cols="12"
            sm="12"
            lg="6"
          >
            <v-card
              color="blue lighten-3"
              height="100%"
              @click="click(article.id)"
            >
              <v-card-title class="headline">{{ article.title }}</v-card-title>

              <v-card-subtitle>{{ article.body }}</v-card-subtitle>

              <v-card-actions> </v-card-actions>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-container>
  </div>
</template>

<script>
export default {
  data: function() {
    return {
      articles: [
        {
          id: 10,
          title: "Title, title, タイトル, たいとる, ﾀｲﾄﾙ",
          body:
            "3分で分かります。高専に任せろ。3分で分かります。高専に任せろ。3分で分かります。高専に任せろ。3分で分かります。高専に任せろ。"
        },
        {
          id: 20,
          title: "Title, title, タイトル, たいとる, ﾀｲﾄﾙ",
          body:
            "3分で分かります。高専に任せろ。3分で分かります。高専に任せろ。3分で分かります。高専に任せろ。3分で分かります。高専に任せろ。"
        },
        {
          id: 30,
          title: "Title, title, タイトル, たいとる, ﾀｲﾄﾙ",
          body:
            "3分で分かります。高専に任せろ。3分で分かります。高専に任せろ。3分で分かります。高専に任せろ。3分で分かります。高専に任せろ。"
        },
        {
          id: 40,
          title: "Title, title, タイトル, たいとる, ﾀｲﾄﾙ",
          body:
            "3分で分かります。高専に任せろ。3分で分かります。高専に任せろ。3分で分かります。高専に任せろ。3分で分かります。高専に任せろ。"
        },
        {
          id: 50,
          title: "Title, title, タイトル, たいとる, ﾀｲﾄﾙ",
          body:
            "3分で分かります。高専に任せろ。3分で分かります。高専に任せろ。3分で分かります。高専に任せろ。3分で分かります。高専に任せろ。"
        }
      ]
    };
  },

  methods: {
    click(id) {
      this.$router.push({ name: "Read-article", params: { id } });
    }
  }
};
</script>
