function computePlaceholder(prop: string | undefined, placeholder: string | undefined) {
  return placeholder || placeholder
  // if (this.placeholder) {
  //   return this.placeholder
  // }
  // let temp = this.$t('field.' + this.fieldName + '_placeholder')
  // if (temp == ('field.' + this.fieldName + '_placeholder')) {
  //   if (this.label) {
  //     return this.$t('field.please_input') + this.label
  //   }
  //   return this.$t('field.please_input') + this.$t('field.' + this.fieldName)
  // } else {
  //   return temp
  // }
}

export default computePlaceholder;