<template>
  <div id="app">
    <PokemonList :pokemonList="pokemonList" />
  </div>
</template>

<script>
import PokemonList from './components/PokemonList.vue';

export default {
  components: {
    PokemonList,
  },
  data() {
    return {
      pokemonList: [],
    };
  },
  mounted() {
    this.fetchPokemonList();
  },
  methods: {
    async fetchPokemonList() {
      try {
        const response = await fetch('https://pokeapi.co/api/v2/pokemon/');
        const data = await response.json();
        this.pokemonList = data.results;
      } catch (error) {
        console.error('Error fetching Pokemon list:', error);
      }
    },
  },
};
</script>

<style>
/* Add any necessary styling here */
</style>
