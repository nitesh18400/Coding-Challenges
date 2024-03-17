<!-- src/components/UrlShortener.vue -->
<template>
    <div class="container">
      <h1 class="mt-5">URL Shortener</h1>
      <div class="input-group mt-3">
        <input type="text" v-model="url" class="form-control" placeholder="Enter URL">
        <button @click="shortenUrl" class="btn btn-primary">Shorten</button>
      </div>
      <p v-if="shortenedUrl" class="mt-3">
        Shortened URL:
        <a :href="shortenedUrl" target="_blank" class="link-primary">{{ shortenedUrl }}</a>
      </p>
    </div>
  </template>̵
  
  <script>
  export default {
    data() {
      return {
        url: '',
        shortenedUrl: ''
      };
    },
    methods: {
      shortenUrl() {
        console.log("hulala", this.url)
        // Make an HTTP POST request to the URL shortening endpoint
      fetch('http://localhost:8080/shorten', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          original_url: this.url
        })
      })
      .then(response => response.json())
      .then(data => {
        // Extract the shortened URL from the response and update the component's data
        console.log("hulala", data);
        this.shortenedUrl = data.short_url;
      })
      .catch(error => {
        // Handle any errors that occur during the request
        console.error('Error:', error);
      });
      }
    }
  };
  </script>
  
  <style scoped>
  /* Add your component-specific styles here */
  </style>
  