<template>
  <div style="display: inline-block;">
    <el-popover triggers="hover" placement="bottom" variant="primary" width="420px" title="在线演示">
      <template #reference>
        <div style="cursor: pointer;" :id="`demo-popover-target-${suffix}`">
          <slot>
            <span :style="`font-weight: bold; color: ${color};`">{{label}}</span>
          </slot>
        </div>
      </template>
      <div style="font-size: 13px; color: #666;">
        在线演示环境中，您具备 <span style="color: red; font-weight: bold">只读</span> 权限，只能体验 Kuboard 的一部分功能。<br/>
      </div>
      <div style="padding: 10px; border: 1px solid #eee; border-radius: 10px; margin: 10px 0px; background-color: #fafafa;">
        <a href="https://demo.kuboard.cn" target="_blank">https://demo.kuboard.cn</a> <br/>
        <div style="width: 60px; display: inline-block;margin-top: 5px;">用&nbsp;户</div>
        demo <br/>
        <div style="width: 60px; display: inline-block">密&nbsp;码</div>
        demo123
      </div>
    </el-popover>
  </div>
</template>

<script>
export default {
  props: {
    suffix: { type: String, required: false, default: '' },
    label: { type: String, required: false, default: 'Demo' },
    color: { type: String, required: false, default: 'auto' },
  },
  data() {
    return {

    }
  },
  computed: {
  },
  components: { },
  mounted () {
  },
  methods: {

  }
}
</script>

<style scoped lang="css">

</style>
