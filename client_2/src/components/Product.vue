<template>
  <div class="product-info" @mouseover="showPrice" @mouseleave="hidePrice">
    <h4 class="name">{{ this.product.productName }}</h4>
    <div class="image-holder">
      <img
        :src="this.product.productImageSource"
        :alt="this.product.productName"
      />
    </div>
    <div class="description-price-buy" :class="showingPrice?'':'hide'" >
      <div class="description">
        <h5>Description</h5>
        <p>
          The Ios Dining Chair captures the essence of resort decorating with
          provincial flair. Featuring a curved arm profile, the Ios is
          constructed from synthetic rattan in easy to co-ordinate natural,
          weathered grey or anthracite finishes. Water resistant cushions
          complete the collection providing the ultimate in dining comfort.
        </p>
      </div>
      <div class="price-buy" :class="showingPrice ? 'show' : 'hide'">
        <p class="price">
          Our price today: <span class="price">
            {{ product.productPrice }}$</span
          >
        </p>
        <div class="buy-decision" v-if="userLogged">
          <div class="select-number">
            <span class="plus">+</span>
            <span class="number">1</span>
            <span class="minus">-</span>
          </div>
          <div class="add-to-cart">
            <button><i class="fas fa-shopping-cart"></i>+</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { computed, ref } from "@vue/reactivity";
import { useStore } from "vuex";
export default {
  name: "Product",
  props: {
    product: Object,
  },
  setup() {
    const store = useStore();
    const userLogged = computed(() => store.getters.getStatus);
    const showingPrice = ref(false);
    const showPrice = () => {
      showingPrice.value = true;
    };
    const hidePrice = () => {
      showingPrice.value = false;
    };
    return { userLogged, showingPrice, showPrice, hidePrice };
  },
};
</script>

<style lang="scss" scoped>
.product-info {
  position: relative;
  width: 450px;
  height: 100%;
  display: inline-flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: stretch;
  margin: 0 calc((100% - 450px) / 2);
  padding: 1rem;
  overflow-y: hidden;
  background-color: $bg-color2;
  // box-shadow: 0px 5px 5px rgba(0, 0, 0, 1), 0px 10px 10px rgba(0, 0, 0, 0.4);
  border: 1px solid rgba(0, 0, 0, 0.1);
  h4 {
    font-size: 1.3rem;
    padding-left: 1rem;
    white-space: normal;
    margin: 0;
  }

  .description-price-buy.hide {
    transform: translateY(250px);
    opacity: 0;
  }
  .description-price-buy {
    transition: all 0.7s;
    position: absolute;
    background: rgba(0, 0, 0, 0.5);
    border-radius: 20px 20px 0 0;
    color: white;
    bottom: 0;
    left: 0;
    width: 100%;
    height: 250px;
    overflow: auto;
    padding: 1rem;
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    div.description {
      h5 {
        font-size: 1rem;
        margin:0;
      }
      p {
        white-space: normal;
        font-size: 0.9rem;
      }
    }
    .price-buy {
      display: flex;
      justify-content: space-between;
      align-items: center;
      p.price {
        margin: 0;
        font-size: 1rem;
        white-space: normal;
        padding-left: 0;
        line-height: 33px;
         span.price {
           
          margin-left: 0.3rem;
        }
      }
      .buy-decision {
        display: flex;
        justify-content: space-around;
        .select-number {
          text-align: center;
          .minus,
          .plus,
          .number {
            --button-side: 30px;
            background: $color1;
            color: white;
            margin-right: 0.2rem;
            display: inline-block;
            width: var(--button-side);
            height: var(--button-side);
            border-radius: calc(var(--button-side) * 0.2);
            line-height: var(--button-side);
          }
        }
        .add-to-cart {
          padding: 0;
          button {
            --button-side: 30px;
            height: var(--button-side);
            border-radius: calc(var(--button-side) * 0.2);
            border: 1px solid rgba(128, 128, 128, 0.4);
            padding: 0 1rem;
            background: white;
            font-size: 1.1rem;
            border-radius: 5px;
          }
        }
      }
    }
  }
  .image-holder {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 80%;
    img {
      max-height: 100%;
      max-width: 100%;
      // height: auto;
    }
  }
}
</style>