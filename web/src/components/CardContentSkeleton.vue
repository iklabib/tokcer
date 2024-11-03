<script setup lang="ts">
import Skeleton from 'primevue/skeleton'
</script>

<template>
    <div class="flex mb-4 w-full">
        <Skeleton size="2.5rem" class="mr-2"></Skeleton>
        <div class="w-full">
            <Skeleton width="10rem" class="mb-2"></Skeleton>
            <Skeleton width="5rem" class="mb-2"></Skeleton>
            <Skeleton height="2rem" class="w-full"></Skeleton>
        </div>
    </div>
</template>