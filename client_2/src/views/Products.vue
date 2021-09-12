<template>
  <div class="main-content"><div class="category-list">
    <Category
      v-for="category in this.allProducts"
      :key="category.id"
      :category="category"
    />
  </div></div>
</template>

<script>
import Category from "../components/Category.vue";
import { useStore} from "vuex";
import { onMounted } from "@vue/runtime-core";
import { computed } from "@vue/reactivity";
export default {
  name: "Home",
  components: {
    Category,
  },
  setup() {
    //access store
    const store = useStore();

    const allProducts = computed(()=>store.getters.allProducts)

    const loadAllProducts = function(){console.log("hello");
    store.dispatch('act_loadAllProducts')}

    //need mounted hook to load all product
    onMounted(loadAllProducts);
    return { allProducts };
  },
};
</script>


<style lang="scss">
.main-content{
  overflow: auto;
   background-image: url(https://media.architecturaldigest.com/photos/5f298951485b66fbbd9021cf/16:9/w_2560%2Cc_limit/IMG-22.jpg);
  // background-color: $bg-color3;
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  background-attachment: fixed;
  width: 100%;
  min-height: 100vh;
}
.category-list{
  margin-left: 450px;  
  overflow: auto;
}
 
</style>
