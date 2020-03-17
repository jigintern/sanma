<template>
  <div class="home">
    <v-container>
      <v-container>
        <div v-show="loading">Loading...</div>
        <v-row dense v-show="!loading">
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

              <v-card-subtitle>{{ article.body | ellipsis }}</v-card-subtitle>

              <v-card-actions>id: {{ article.user.id }}</v-card-actions>
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
      loading: true,
      articles: []
      /*
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
      */
    };
  },

  filters: {
    ellipsis(value) {
      if (!value) return "";
      if (value.length > 10) {
        return value.slice(0, 100) + "...";
      }
      return value;
    }
  },

  mounted: function() {
    this.$axios.get("https://qiita.com/api/v2/items").then(res => {
      console.log(res);
      this.articles = res.data;
      this.loading = false;
    });
  },

  methods: {
    click(id) {
      this.$router.push({ name: "Read-article", params: { id } });
    }
  }
};
</script>
