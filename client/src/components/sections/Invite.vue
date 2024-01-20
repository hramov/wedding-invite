<script setup lang="ts">
</script>

<template>
  <div class="section section_center section_padding" id="section-invite">
    <img alt="flowers" style="position:absolute; top: 0; right: 0; width: 50%; max-width: 500px"
      src="../../assets/img/header.png" />
    <div class="section-body">

      <div class="section__invite">
        <!-- Имена -->
        <h2>
          <span>
            Анастасия
            <span>&amp;</span>

            Сергей </span>
        </h2>

        <h3>Приглашают на свою свадьбу</h3>
        <h2 style="color: black">02.06.2024</h2>

        <!-- Текст -->
        <div class="section__invite-text js-emoji-included">
          <p>Дорогие друзья и родные!<br />
            <br />
            Приглашаем разделить с нами радость особенного для нас события и стать свидетелями начала нашей семейной
            жизни!
            <br />
            <br />
            <!--						Мы будем очень рады, если вы разделите этот счастливый день с нами. <br />-->
            <!--						<br />-->
            <!--						Увидимся на нашей свадьбе!-->
          </p>
        </div>

        <!-- Видео-обращение -->

        <!-- Жених и невеста -->
        <div class="section__invite-newlyweds">
          <div class="item">
            <img src="../../assets/img/nastya3.jpg" alt="Анастасия" />
          </div>

          <div class="item">
            <img src="../../assets/img/sergey.jpg" alt="Сергей" />
          </div>
        </div>

        <!-- Жених и невеста -->
        <!--				<div class="section__invite-logo"></div>-->

        <div class="section__first-countdown" id="simple-countdown">

          <div class="simple-countdown">
            <div class="simple-countdown-value oopacity-0" style="width: 23%">
              <span class="simple-countdown-days">00</span>
              Дней
            </div>
            <div class="simple-countdown-value oopacity-0" style="width: 23%">
              <span class="simple-countdown-hours">00</span>
              Часов
            </div>
            <div class="simple-countdown-value oopacity-0" style="width: 23%">
              <span class="simple-countdown-minutes">00</span>
              Минут
            </div>
            <div class="simple-countdown-value oopacity-0" style="width: 23%">
              <span class="simple-countdown-seconds">00</span>
              Секунд
            </div>
          </div>
        </div>
      </div>

    </div>
  </div>
</template>

<style scoped>
.section__first-countdown .simple-countdown {
  /** display: grid;
    grid-template-columns: repeat(4,1fr);
    grid-template-rows: 1fr;
    grid-column-gap: 24px;
    width: 600px;
    height: 132px;
    grid-row-gap: 0;
    */
  display: flex;
  width: 100%;
  justify-content: space-between;
  height: 132px;
  flex-wrap: wrap;
  margin-top: 70px;
}

.section__first-countdown .simple-countdown-value {
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  background: rgba(17, 132, 95, 0.1);
  backdrop-filter: blur(4px);
  transition: opacity .15s ease-in-out;
  border-radius: 16px;
  font-size: 14px;
  text-transform: uppercase;
}

.section__first-countdown .simple-countdown-value span {
  display: block;
  font-size: 42px;
  line-height: 1;
  margin-bottom: 4px;
}
</style>