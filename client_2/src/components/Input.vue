<template>
  <input :placeholder="placeholder" :type="type" :id="id" />
</template>

<script>
export default {
  name: "Input",
  props: {
    placeholder: String,
    type: String,
    id: String,
  },
};
</script>

<style lang="scss" scoped>
input {
  display: block;
  width: 100%;
  &::placeholder {
    font-style: italic;
  }
  display: block;
  width: 100%;
  font-size: 1.3rem;
  padding: 1rem 2rem;
  border: none;
  border-radius: 27px;
  box-shadow: 0 5px 10px rgba(0, 0, 0, 0.3);
  &:focus-visible {
    outline: none;
    &::placeholder {
      color: white;
    }
  }
}
</style>