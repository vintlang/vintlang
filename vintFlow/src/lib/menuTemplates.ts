import { app, BrowserWindow, ipcMain, dialog, Menu, MenuItem, MenuItemConstructorOptions } from 'electron';
import fs from 'fs'
import os from 'os'
import path from 'path'
import Store from 'electron-store';


const isMac = process.platform === 'darwin'

let mainWindow: BrowserWindow;

export const get_files = (path: string) => {
    const files = fs.readdirSync(path, {
      recursive: true,
      withFileTypes: true
    })
    // return files.map(file => ({...file, is_dir: file.isDirectory(), is_symlink: file.isSymbolicLink()})).filter(file => file.name != '.git')
    return files.map(file => ({...file, is_dir: file.isDirectory(),})).filter(file => file.name != '.git')
}

export const menuTemplate = [
    // { role: 'appMenu' }
    ...(isMac
      ? [{
          label: app.name,
          submenu: [
            { role: 'about' },
            { label: 'Check For Update...' },
            { type: 'separator' },
            { label: 'Settings',
              submenu: [
                {
                  label: "Profiles Default",
                  submenu: [
                    {
                      label: 'Default',
                      type: 'checkbox',
                      checked: true
                    },
                    {
                      type: 'separator'
                    },
                    {
                      label: 'Show Profile Content',
                      type: 'checkbox',
                      checked: false
                    },
                    {
                      type: 'separator'
                    },
                    {
                      label: 'Show Profile',
                      type: 'checkbox',
                      checked: false
                    },
                    {
                      label: 'Delete Profile',
                      type: 'checkbox',
                      checked: false,
                      enabled: false
                    },
                    {
                      type: 'separator'
                    },
                    {
                      label: 'Export',
                      type: 'checkbox',
                      checked: false
                    },
                    {
                      label: 'Import',
                      type: 'checkbox',
                      checked: false,
                    },
                  ]
                },
                {label: 'Settings'},
                {label: 'Extensions'},
                {label: 'Keyboard Shortcuts'},
                {label: 'Configure User Snippets'},
                {label: 'User Tasks'},
                {
                  label: 'Theme',
                  submenu: [
                    {label: 'Color Theme'},
                    {label: 'File Icon Theme'},
                    {label: 'Product Icon Theme'}
                  ]
                },
                {type: 'separator'},
                {label: 'Online Services Settings'},
                {type: 'separator'},
                {label: 'Backup and Sync Settings'},
              ]
            },
            { type: 'separator' },
            { role: 'services' },
            { type: 'separator' },
            { role: 'hide' },
            { role: 'hideOthers' },
            { role: 'unhide' },
            { type: 'separator' },
            { role: 'quit' }
          ]
        }]
      : []),
    // { role: 'fileMenu' }
    {
      label: 'File',
      submenu: [
        {label: 'New Vint File'},
        {label: 'New File'},
        {label: 'New Window'},
        {type: 'separator'},
        {label: 'Open...'},
        {label: 'Open Folder...', click: async () => {
          const folder = await dialog.showOpenDialog(mainWindow, {properties: ['openDirectory']})
          let structure = undefined;
          if (!folder.canceled) {
            console.log("folder", folder.filePaths[0]);
            const tree = get_files(folder.filePaths[0])
            structure = {
              // name: path.dirname(folder.filePaths[0]),
              name: folder.filePaths[0],
              root: folder.filePaths[0],
              tree,
            }  
            // @ts-ignore
            store.set(SELECTED_FOLDER_STORE_NAME, structure)
            // ipcMain.emit('new-folder-opened')
            mainWindow.webContents.send('new-folder-opened')
          }
        }},
        {label: 'Open Workspace From File...'},
        {label: 'Open Recent', submenu: [
          {label: 'Recent File'}
        ]},
        {type: 'separator'},
        {label: 'Add Folder to Workspace'},
        {label: 'Save Workspace As'},
        {label: 'Duplicate Workspace'},
        {type: 'separator'},
        {label: 'Save'},
        {label: 'Save As...'},
        {label: 'Save All'},
        {type: 'separator'},
        {label: 'Share', submenu: [
          {label: "Export Profile..."},
          {label: "Import Profile..."}
        ]},
        {type: 'separator'},
        {label: 'Autosave', type: 'checkbox', checked: true},
        {type: 'separator',},
        {label: 'Revert File'},
        isMac ? {label: 'Close Editor', role: 'close' } : {label: 'Close Editor', role: 'quit' },
        {label: 'Close Folder', role: isMac ? 'close' : 'quit'},
        isMac ? { role: 'close' } : { role: 'quit' }
      ]
    },
    {
      label: 'Edit',
      submenu: [
        { role: 'undo' },
        { role: 'redo' },
        { type: 'separator' },
        { role: 'cut' },
        { role: 'copy' },
        { role: 'paste' },
        ...(isMac
          ? [
              { role: 'pasteAndMatchStyle' },
              { role: 'delete' },
              { role: 'selectAll' },
              { type: 'separator' },
              {
                label: 'Speech',
                submenu: [
                  { role: 'startSpeaking' },
                  { role: 'stopSpeaking' }
                ]
              }
            ]
          : [
              { role: 'delete' },
              { type: 'separator' },
              { role: 'selectAll' }
            ])
      ]
    },
    // { role: 'editMenu' }
    {
      label: 'Selection',
      submenu: [
        { role: 'selectAll' },
        { label: 'Expand Selection' },
        { label: 'Shrink Selection' },
        { type: 'separator' },
        { label: 'Copy Line Up' },
        { label: 'Copy Line Down' },
        { label: 'Move Line Up' },
        { label: 'Move Line Down' },
        { label: 'Duplicate Selection' },
        { type: 'separator' },
        { label: 'Add Cursor Above' },
        { label: 'Add Cursor Below' },
        { label: 'Add Cursor to Line Ends' },
        { label: 'Add Next Occurrence' },
        { label: 'Add Previous Occurrence' },
        { label: 'Select All Occurrence' },
        { type: 'separator' },
        { label: 'Switch '+ (isMac ? 'Cmd+Click' : "Control+Click") +' to Multi-Cursor' },
        { label: 'Column Selection Mode' },
      ]
    },
    {
      label: 'View',
      submenu: [
        { label: 'Command Pallete' },
        { label: 'Open View' },
        { type: 'separator' },
        { label: 'Apperance', submenu: [
          {label: "Full Screen"},
          {label: "Zen Mode"},
          {label: "Center Layout"},
        ] },
        { label: 'Editor Layout', submenu: [
          {label: "Split Up"},
          {label: "Split Down"},
          {label: "Split Left"},
          {label: "Split Right"},
          { type: 'separator' },
          {label: "Split In Group"},
          { type: 'separator' },
          {label: "Move Editor into New Window"},
          {label: "Copy Editor into New Window"},
          { type: 'separator' },
          {label: "Single"},
          {label: "Two Columns"},
          {label: "Three Columns"},
          {label: "Two Rows"},
          {label: "Grid (2x2)"},
          {label: "Two Rows Right"},
          {label: "Two Columns Bottom"},
        ] },
        { type: 'separator' },
        { label: 'Explorer' },
        { label: 'Search' },
        { label: 'Source Control' },
        { label: 'Run' },
        { label: 'Extensions' },
        { type: 'separator' },
        { label: 'Problems' },
        { label: 'Output' },
        { label: 'Debug Console' },
        { label: 'Teminal' },
        { type: 'separator' },
        { label: 'Word Wrap' },
      ]
    },
    // { role: 'windowMenu' }
    {
      label: 'Go',
      submenu: [
        { label: 'Back' },
        { label: 'Forward', enabled: false },
        { label: 'Last Edit Location' },
        { label: 'Switch Editor', submenu: [
          {label: "Next Editor"},
          {label: "Previous Editor"},
          {type: 'separator'},
          {label: "Next Used Editor"},
          {label: "Previous Used Editor"},
          {type: 'separator'},
          {label: "Next Editor in Group"},
          {label: "Previous Editor in Group"},
          {type: 'separator'},
          {label: "Next Used Editor in Group"},
          {label: "Previous Used Editor in Group"},
        ] },
        { label: 'Switch Group', submenu: [
          {label: 'Group 1'},
          {label: 'Group 2'},
          {label: 'Group 3', enabled: false},
          {label: 'Group 4', enabled: false},
          {label: 'Group 5', enabled: false},
          {type: 'separator'},
          {label: 'Next Group', enabled: false},
          {label: 'Previous Group', enabled: false},
          {type: 'separator'},
          {label: 'Group Left', enabled: false},
          {label: 'Group Right', enabled: false},
          {label: 'Group Above', enabled: false},
          {label: 'Group Below', enabled: false},
        ] },
        {type: 'separator'},
        { label: 'Go to File' },
        { label: 'Go to Symbol in Workspace' },
        {type: 'separator'},
        { label: 'Go to Symbol in Editor' },
        { label: 'Go to Definition' },
        { label: 'Go to Declaration' },
        { label: 'Go to Type Definition' },
        { label: 'Go to Implementations' },
        { label: 'Go to References' },
        {type: 'separator'},
        { label: 'Go to Line/Column' },
        { label: 'Go to Bracket' },
        {type: 'separator'},
        { label: 'Next Problem' },
        { label: 'Previous Problem' },
        {type: 'separator'},
        { label: 'Next Change' },
        { label: 'Previous Change' },      
      ]
    },
    {
      label: 'Run',
      submenu: [
        { label: 'Start Debugging' },
        { label: 'Run Without Debugging' },
        { label: 'Stop Debugging', enabled: false },
        { label: 'Restart Debugging', enabled: false },
        {type: 'separator'},
        { label: 'Open Configuration', enabled: false },
        { label: 'Add Configuration', enabled: true },
        {type: 'separator'},
        { label: 'Step Over', enabled: false },
        { label: 'Step Into', enabled: false },
        { label: 'Step Out', enabled: false },
        { label: 'Continue', enabled: false },
        {type: 'separator'},
        { label: 'Toggle Breakpoint', },
        { label: 'New Breakpoint', },
        { role: 'zoom', submenu: [
          {label: 'Conditional Breakpoint'},
          {label: 'Edit Breakpoint'},
          {label: 'Inline Breakpoint'},
          {label: 'Function Breakpoint...'},
          {label: 'Logpoint...'},
        ] },
        {type: 'separator'},
        { label: 'Enable All Breakpoints', },
        { label: 'Disable All Breakpoints', },
        { label: 'Remove All Breakpoints', },
        {type: 'separator'},
        { label: 'Install Additional Debugger...', },
      ]
    },
    {
      label: 'Terminal',
      submenu: [
        { label: 'New Teminal' },
        { label: 'Split Teminal', enabled: false },
        {type: 'separator'},
        { label: 'Run Task', },
        { label: 'Run Build Task', },
        { label: 'Run Active Task', },
        { label: 'Run Selected Task', },
        {type: 'separator'},
        { label: 'Show Running Task', enabled: false },
        { label: 'Restart Running Task', enabled: false },
        { label: 'Terminate Task', enabled: false },
        {type: 'separator'},
        { label: 'Configure Task...', enabled: false },
        { label: 'Configure Default Build Task...', enabled: false },
        { role: 'zoom' },
      ]
    },
    {
      label: 'Window',
      submenu: [
        { role: 'minimize' },
        { role: 'zoom' },
        {label: 'Tile Window to Left of Screen'},
        {label: 'Tile Window to Right of Screen'},
        {label: 'Replace Tiled Window'},
        {type: 'separator'},
        {label: 'Remove Window From Set', enabled: false},
        {type: 'separator'},
        {label: 'Switch Window'},
        {type: 'separator'},
        {label: 'Bring All Front'},
      ]
    },
    {
      role: 'help',
      submenu: [
        { label: 'Welcome',},
        { label: 'Show All Commands',},
        { label: 'Documentation',},
        { label: 'Editor Playground',},
        { label: 'Show Release Notes',},
        {type: 'separator'},
        { label: 'Keyboard Shortcuts Reference',},
        { label: 'Video Tutorials',},
        { label: 'Tips and Tricks',},
        {type: 'separator'},
        { label: 'Join Us on YouTube',},
        { label: 'Switch Feature Request',},
        { label: 'Report Issue',},
        {type: 'separator'},
        { label: 'View License',},
        { label: 'Privacy Statement',},
        {type: 'separator'},
        { label: 'Toggle Developer Tools',},
        { label: 'Open Process Explorer',},
      ]
    }
  ] as unknown as MenuItem[]
  
