<template>
  <div class="category">
    <h4>{{ category.categoryName }}</h4>
    <div class="product-list" :style="{ left: leftPosition + 'px' }">
      <Product
        v-for="product in category.productArray"
        :key="product.productID"
        :product="product"
      />
    </div>
    <div class="wipe-handler">
      <div class="handler" id="go-left" @click="wipe">
        <i class="fas fa-sort-up"></i>
      </div>
      <div class="handler" id="go-right" @click="wipe">
        <i class="fas fa-sort-up"></i>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from "@vue/reactivity";
import Product from "./Product.vue";
export default {
  name: "Category",
  components: { Product },
  props: {
    category: Object,
  },
  setup() {
    const leftPosition = ref(0);
    const wipe = (event) => {
      const direction = event.target.id;
      const distanceToGo = event.target.parentNode.parentNode.offsetWidth;
      let step = Array.prototype.slice.call(
        event.target.parentNode.previousSibling.childNodes
      );
      console.log(step);
      step = step.filter((e) => e.nodeName != "#text");
      const stepLimit = step.length;
      const distanceLimit = (stepLimit - 1) * distanceToGo;
      console.log(leftPosition.value, distanceToGo, stepLimit, distanceLimit);
      if (direction == "go-left") {
        if (leftPosition.value < 0) {
          leftPosition.value += distanceToGo;
        }
      } else {
        if (leftPosition.value > -distanceLimit) {
          leftPosition.value -= distanceToGo;
        }
      }
    };

    return { leftPosition, wipe };
  },
};
</script>

<style lang="scss" scoped>
.category {
  position: relative;
  --width: calc((100vw - 450px) * 0.7);
  --height: calc(var(--width) * 0.768);
  width: var(--width);
  height: var(--height);
  padding: 1rem 0;
  margin: auto;
  margin-top: calc((100vh - var(--height)) / 2);
  background-color: rgba(255, 255, 255, 0.6);
  border-radius: 15px;
  overflow-x: hidden;
  // box-shadow: 0 40px 0px rgb(184, 182, 182,0.6);
  color: gray;
  display: flex;
  flex-direction: column;
  justify-content: space-evenly;
  align-items: center;

  margin-bottom: calc(100vh - var(--height));
  h4 {
    // height: 40px;
    padding: 0.5rem 1rem;
    font-size: 1.4rem;
    width:100%;
    background-color: $bg-color4;
    margin: 0;
    text-align: center;
  }
  .product-list {
    white-space: nowrap;
    overflow-x: visible;
    height: calc(var(--height) * 0.7);
    width: 100%;
    position: relative;
    transition: left 0.5s;
  }
  .wipe-handler {
    pointer-events: none;
    position: absolute;
    width: 450px;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
    display: flex;
    justify-content: space-between;
    div.handler {
      pointer-events: fill;
      font-size: 3rem;
      background: rgba(0, 0, 0, 0.15);
      color: white;
      display: flex;
      justify-content: center;
      align-items: center;
      width: 70px;
      height: 70px;
      border-radius: 35px;
      cursor: pointer;
      i {
        height: 30px;
        pointer-events: none;
      }
    }
    div.handler#go-left {
      transform: rotate(-90deg);
    }
    div.handler#go-right {
      transform: rotate(90deg);
    }
  }
}
</style>