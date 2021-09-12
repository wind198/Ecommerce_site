<template>
  <button :style="{ backgroundColor: color }" >{{ text }}</button>
</template>

<script>
export default {
  name: "Button",
  props: { text: String, color: String },
};
</script>

<style lang="scss" scoped>
button {
  display: block;
  width: 100%;
  font-size: 1.4rem;
  padding: 1rem 1rem;
  text-transform: capitalize;
  letter-spacing: 0.1rem;
  border-radius: 30px;
  box-shadow: 0 5px 10px rgba(0, 0, 0, 0.3);
  &:focus-visible {
    font-weight: bold;
  }
}
</style>