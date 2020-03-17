<template>
  <div class="home">
    <v-container>
      <v-container>
        <div v-show="loading">Loading...</div>
        <v-row dense v-show="!loading">
          <v-col v-for="(article, index) in articles" :key="index" cols="12" sm="12" lg="6">
            <v-card color="blue lighten-3" height="100%" @click="click(article.id)">
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
    this.$axios.get("https://sanma.intern.jigd.info/api/articles").then(res => {
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
