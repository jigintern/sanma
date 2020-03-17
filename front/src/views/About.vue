<template>
  <v-app>
    <v-card width="800px" class="mx-auto mt-5">
      <v-card-title>
        <h1 class="display-3 font-weight-black">sanma</h1>
        <v-card-text>
          <div class="display-1">
            みんなで作る
            <br />3分で知見を得られるTechメディア
          </div>
          <br />
          <div class="headline font-weight-black">#こんな人に向いてる</div>
          <ul class="headline">
            <li>忙しい人</li>
            <li>勉強の仕方が分からん人</li>
            <li>短い時間で知見獲得したい人</li>
          </ul>
          <br />
          <div class="headline">自分のブログに適用して</div>

          <v-card-actions>
            <v-btn class="info text-none" large @click="download">MarkDown Download</v-btn>
          </v-card-actions>
        </v-card-text>
      </v-card-title>
    </v-card>
  </v-app>
</template>

<script>
export default {
  name: "About",
  methods: {
    download: function() {
      console.log("download");
      alert("ダウンロードできませんｗ");
    }
  }
};
</script>
