<script setup lang="ts">
import {onMounted, ref} from "vue";
import axios from 'axios';

const scrollToFormStart = () =>  {
  const el = document.getElementById('section-poll');

  if (el) {
    el.scrollIntoView({ behavior: 'smooth'});
  }
}

const companions = ref([]);

const getCompanions = async () => {
  try {
    const resp = await axios.get(import.meta.env.VITE_API_URL + '/companions');
    if (resp.status < 400) {
      companions.value = resp.data;
    }
  } catch (_err: unknown) {
    const err = _err as Error
    console.log(err.message)
  }
}

onMounted(async () => {
  await getCompanions();
});

const poll = ref({
  name: '',
  email: '',
  presence: '',
  drink: [] as string[],
  food: [] as string[],
  music: '',
  transfer: '',
  accommodation: '',
  companion: null,
});

const clear = () => {
  poll.value = {
    name: '',
    email: '',
    presence: '',
    drink: [],
    food: [],
    music: '',
    transfer: '',
    accommodation: '',
    companion: null,
  };
}

const errors = ref<{ field: string, message: string}[]>([]);

const successMessage = ref('');

const errorMessage = ref('');

const validate = (): any[] => {
  const errors = [];

  if (!poll.value.name) {
    errors.push({
      field: 'name',
      message: 'Инкогнито из города N должен представиться',
    });
  }

  if (!poll.value.email) {
    errors.push({
      field: 'email',
      message: 'Сюда будут приходить новости',
    });
  }

  if (!poll.value.presence) {
    errors.push({
      field: 'presence',
      message: 'Пожалуйста, заполните присутствие',
    });
  }

  if (poll.value.presence === 'yes') {
    if (!poll.value.drink.length) {
      errors.push({
        field: 'alcohol',
        message: 'Пожалуйста, дайте нам знать, что вы предпочитаете пить'
      });
    }

    if (!poll.value.food.length) {
      errors.push({
        field: 'food',
        message: 'Пожалуйста, дайте нам знать, что вы предпочитаете из еды'
      });
    }

    if (!poll.value.music) {
      errors.push({
        field: 'music',
        message: 'Ну хоть одну песенку :-)'
      });
    }

    if (!poll.value.transfer) {
      errors.push({
        field: 'transfer',
        message: 'Пожалуйста, дайте нам знать, сможете ли вы добраться сами или понадобится помощь'
      });
    }

    if (!poll.value.accommodation) {
      errors.push({
        field: 'accommodation',
        message: 'Пожалуйста, дайте нам знать, сможете ли вы разместиться сами или понадобится помощь'
      });
    }
  }

  if (errors.length > 0) {
    scrollToFormStart()
  }

  return errors;
}

const submit = async () => {
  if (poll.value.drink.includes('Другое')) {
    poll.value.drink = poll.value.drink.filter((item: string) => item !== 'Другое');
    poll.value.drink.push(otherDrink.value)
  }

  errors.value = validate();

  console.log(errors.value);
  if (errors.value.length) {
    return;
  }

  errorMessage.value = '';
  successMessage.value = '';

  const error = await sendToServer();

  if (error === null) {
    successMessage.value = 'Спасибо! Ваши ответы отправлены'
    clear();
  } else {
    errorMessage.value = 'Что-то пошло не так... Попробуйте еще раз';
  }
  scrollToFormStart();
}

const sendToServer = async () => {
  try {
    await axios.post(import.meta.env.VITE_API_URL + '/add-guest', poll.value);
    await getCompanions();
    return null
  } catch (err) {
    return err;
  }
}

const alcohol: Array<{title: string, value: string}> = [
  {
    title: 'Вино белое',
    value: 'Вино белое'
  },
  {
    title: 'Вино красное',
    value: 'Вино красное'
  },
  {
    title: 'Шампанское',
    value: 'Шампанское'
  },
  {
    title: 'Водка',
    value: 'Водка'
  },
  {
    title: 'Виски',
    value: 'Виски'
  },
  {
    title: 'Коньяк',
    value: 'Коньяк'
  },
  {
    title: 'Безалкогольные напитки',
    value: 'Безалкогольные напитки'
  },
  {
    title: 'Другое',
    value: 'Другое'
  },
];

const otherDrink = ref<string>('');

const food: Array<{title: string, value: string}> = [
  {
    title: 'Рыба',
    value: 'Рыба'
  },
  {
    title: 'Мясо',
    value: 'Мясо'
  },
  {
    title: 'Птица',
    value: 'Птица'
  },
  {
    title: 'Веган',
    value: 'Веган'
  },
]
</script>

<template>
  <h3 style="color: #11845f">{{ successMessage }}</h3>
  <h3 style="color: darkred">{{ errorMessage }}</h3>

  <form @submit.prevent="submit">
    <p style="color: darkred" v-if="errors.filter((item: any) => item.field === 'name').length">{{ errors.find((item: any) => item.field === 'name')?.message }}</p>
    <v-text-field
        v-model="poll.name"
        :counter="10"
        label="Фамилия и имя"
    ></v-text-field>

    <p style="color: darkred" v-if="errors.filter((item: any) => item.field === 'email').length">{{ errors.find((item: any) => item.field === 'email')?.message }}</p>
    <v-text-field
        v-model="poll.email"
        type="email"
        label="E-mail"
    ></v-text-field>

    <v-radio-group v-model="poll.presence">
      <p>Присутствие</p>
      <p style="color: darkred" v-if="errors.filter((item: any) => item.field === 'presence').length">{{ errors.find((item: any) => item.field === 'presence')?.message }}</p>
      <v-radio value="yes">
        <template v-slot:label>
          <div>Я с удовольствием приду</div>
        </template>
      </v-radio>
      <v-radio value="no">
        <template v-slot:label>
          <div>К сожалению, не смогу присутствовать</div>
        </template>
      </v-radio>
    </v-radio-group>

    <div v-if="poll.presence === 'yes'">
      <div>
        <p>Что предпочитаете из напитков?</p>
        <p style="color: darkred" v-if="errors.filter((item: any) => item.field === 'drink').length">{{ errors.find((item: any) => item.field === 'drink')?.message }}</p>
        <v-checkbox
            style="margin: 0; padding: 0; height: 50px"
            v-for="item in alcohol" :key="item.title"
            v-model="poll.drink"
            :label="item.title"
            :value="item.value"
        ></v-checkbox>
        <v-text-field
            v-if="poll.drink.includes('Другое' as never)"
            v-model="otherDrink"
            label="Ваш вариант"
        ></v-text-field>
      </div>

      <div>
        <p>Ваши предпочтения в еде?</p>
        <p style="color: darkred" v-if="errors.filter((item: any) => item.field === 'food').length">{{ errors.find((item: any) => item.field === 'food')?.message }}</p>
        <v-checkbox
            style="margin: 0; padding: 0; height: 50px"
            v-for="item in food" :key="item.title"
            v-model="poll.food"
            :label="item.title"
            :value="item.value"
        ></v-checkbox>
      </div>
      <br />
      <div>
        <p>Напишите несколько музыкальных композиций, которые Вы бы хотели услышать</p>
        <p style="color: darkred" v-if="errors.filter((item: any) => item.field === 'music').length">{{ errors.find((item: any) => item.field === 'music')?.message }}</p>
        <v-textarea v-model="poll.music" placeholder="1. Виски, кола, королева танцпола"></v-textarea>
      </div>

      <v-radio-group v-model="poll.transfer">
        <p>Нужен ли вам трансфер?</p>
        <p style="color: darkred" v-if="errors.filter((item: any) => item.field === 'transfer').length">{{ errors.find((item: any) => item.field === 'transfer')?.message }}</p>
        <v-radio value="Да, до места торжества и обратно">
          <template v-slot:label>
            <div>Да, до места торжества и обратно</div>
          </template>
        </v-radio>
        <v-radio value="Да, только до места торжества">
          <template v-slot:label>
            <div>Да, только до места торжества</div>
          </template>
        </v-radio>
        <v-radio value="Нет">
          <template v-slot:label>
            <div>Нет, доберусь сам(а)</div>
          </template>
        </v-radio>
      </v-radio-group>

      <v-radio-group v-model="poll.accommodation">
        <p>Требуется ли вам помощь с размещением?</p>
        <p style="color: darkred" v-if="errors.filter((item: any) => item.field === 'accommodation').length">{{ errors.find((item: any) => item.field === 'accommodation')!.message }}</p>
        <v-radio value="Да">
          <template v-slot:label>
            <div>Да, мне нужен номер</div>
          </template>
        </v-radio>
        <v-radio value="Нет">
          <template v-slot:label>
            <div>Нет, я справлюсь сам(а)</div>
          </template>
        </v-radio>
      </v-radio-group>

      <div>
        <p>Если вы будете не один, выберите своего спутника ниже</p>
        <small>(он должен уже заполнить эту форму)</small>
        <v-select
            v-model="poll.companion"
            label="Выберите спутника"
            :items="companions"
            item-title="name"
            item-value="id"
        ></v-select>
      </div>
    </div>

    <v-btn
        class="me-4"
        color="#11845f"
        type="submit"
    >
      Отправить
    </v-btn>
    <v-btn @click="clear" color="#34495e">
      Очистить
    </v-btn>
  </form>
</template>