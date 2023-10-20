

routes.push({ path: '/', name: 'home', component: {
  template: '#page-home',
  data() { return {
    i18n: this.$root.i18n,
    changeDetails: {},
    motd: "",
  }},
  created() {
    this.loadChanges();
  },
  watch: {
  },
  methods: {
    async loadChanges() {
      try {
        const response = await this.$root.createApi().get('motd.md?v=' + Date.now());
        if (response.data) {
          this.motd = response.data;
        }
      } catch (error) {
        this.$root.showError(error);
      }
    },
  }
}});
