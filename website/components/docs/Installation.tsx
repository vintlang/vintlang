import React from 'react'
import SectionHeader from './SectionHeader'
import MarkdownRenderer from '../MarkdownRender'
import { getMarkdownContent } from '@/lib/utils'

const Installation = () => {
    const markdown = getMarkdownContent("README.md")

  return (
    <div className="" id='installation'>
        <SectionHeader title="Installation"/>

        <MarkdownRenderer markdown={markdown}/>

    </div>
  )
}

export default Installation